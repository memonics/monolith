package health

import (
	"context"
	"log"
	"time"
)

type AbsorbedStateController interface {
	SetAbsorbedMode(active bool)
}

type MonolithSupervisor struct {
	monitor    *DomainMonitor
	controller AbsorbedStateController
}

func NewMonolithSupervisor(m *DomainMonitor, c AbsorbedStateController) *MonolithSupervisor {
	return &MonolithSupervisor{monitor: m, controller: c}
}

// WatchAndGovern orchestrates local Tier 2/3 state machine mutations on backoff schedules
func (s *MonolithSupervisor) WatchAndGovern(ctx context.Context, checkInterval time.Duration) {
	ticker := time.NewTicker(checkInterval)
	go func() {
		defer ticker.Stop()
		consecutiveFailures := 0

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s.monitor.mu.RLock()
				healthy := s.monitor.isHealthy
				s.monitor.mu.RUnlock()

				if !healthy {
					consecutiveFailures++
					log.Printf("[SUPERVISOR] Ring 1 health check failed. Failure count: %d", consecutiveFailures)
					
					// Tier 2 degradation threshold reached
					if consecutiveFailures >= 2 && consecutiveFailures < 4 {
						log.Println("[SUPERVISOR] Entering Tier 2: Degraded State. Activating transparent fallback storage paths.")
					}
					
					// Tier 3 unrecoverable dependency boundary hit -> Enforce Write Rejection
					if consecutiveFailures >= 4 {
						log.Println("[SUPERVISOR] Entering Tier 3: Critical System Fault. Mutating coordinator state to WRITE-REJECTION mode.")
						s.controller.SetAbsorbedMode(true)
					}
				} else {
					if consecutiveFailures > 0 {
						log.Println("[SUPERVISOR] Ring 1 recovered baseline operations. Resetting service state machine coordinates.")
						s.controller.SetAbsorbedMode(false)
					}
					consecutiveFailures = 0
				}
			}
		}
	}()
}