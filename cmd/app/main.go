package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	ring0Public "github.com/bdra-io/monolith/internal/ring0/public"
	"github.com/bdra-io/monolith/internal/ring1/health"
	"github.com/bdra-io/monolith/internal/ring1/public"
	
	// Mount Ring 2 components into the global compiler graph
	ring2Health "github.com/bdra-io/monolith/internal/ring2/health"
	ring2Public "github.com/bdra-io/monolith/internal/ring2/public"
	
	_ "github.com/lib/pq"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := sql.Open("postgres", "postgres://user:pass@localhost:5432/db?sslmode=disable")
	if err != nil {
		log.Fatalf("Fatal: Database attachment connection failure: %v", err)
	}

	// 1. Compile Ring 0 (Core Identity)
	identityService := ring0Public.NewUserService()

	// 2. Compile Ring 1 (Transactional Core)
	idGenerator := public.NewCryptoIDGenerator()
	cacheStore := public.NewRedisCacheStore()
	orderRepository := public.NewPostgresOrderRepository(db)
	orderService := public.NewOrderService(orderRepository, idGenerator, cacheStore, identityService)

	// 3. Compile Ring 2 (Operational Intelligence)
	auditService := ring2Public.NewAuditLogService()

	// 4. Launch Asynchronous Out-of-Band Health Monitors
	ring1Monitor := health.NewDomainMonitor()
	ring1Monitor.StartWatchLoop(ctx, 10*time.Second)

	ring2Monitor := ring2Health.NewDomainMonitor()
	ring2Monitor.StartWatchLoop(ctx, 10*time.Second)

	// 5. Build Local Process Resilience Supervisors
	resilienceSupervisor := health.NewMonolithSupervisor(ring1Monitor, orderService)
	resilienceSupervisor.WatchAndGovern(ctx, 3*time.Second)

	// 6. Establish Global Mux Router Mapping
	mux := http.NewServeMux()
	
	// Expose Health Checks for all active rings
	mux.Handle("/ring1/health", ring1Monitor)
	mux.Handle("/ring2/health", ring2Monitor)
	
	// Delegate Route Configuration to Ring HTTP Adapters
	orderHTTPAdapter := public.NewHTTPOrderAdapter(orderService)
	orderHTTPAdapter.RegisterRoutes(mux)

	log.Println("BDRA Lite Fully-Rigged Monolith running seamlessly on port :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Fatal structural server crash: %v", err)
	}
}