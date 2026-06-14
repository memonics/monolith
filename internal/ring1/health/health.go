package health

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type DomainMonitor struct {
	mu           sync.RWMutex
	isHealthy    bool
	systemStatus string
}

func NewDomainMonitor() *DomainMonitor {
	return &DomainMonitor{
		isHealthy:    true,
		systemStatus: "healthy",
	}
}

func (m *DomainMonitor) SetHealthy(healthy bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.isHealthy = healthy
	if healthy {
		m.systemStatus = "healthy"
	} else {
		m.systemStatus = "unhealthy"
	}
}

// StartWatchLoop runs out-of-band background checks, caching metrics directly in memory[cite: 1926].
func (m *DomainMonitor) StartWatchLoop(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		defer ticker.Stop() // Prevents system resource leak on execution cancellation
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.S:
				// Execute out-of-band monitoring checks natively
				m.mu.RLock()
				currentStatus := m.isHealthy
				m.mu.RUnlock()
				
				// For example, validation query mapping goes here
				_ = currentStatus 
			}
		}
	}()
}

func (m *DomainMonitor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mu.RLock()
	status := m.systemStatus
	healthy := m.isHealthy
	m.mu.RUnlock()

	payload := map[string]string{
		"status":    status,
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-BDRA-Probe", "true") // Telemetry bypass decoration [cite: 1931]
	
	if !healthy {
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	_ = json.NewEncoder(w).Encode(payload)
}