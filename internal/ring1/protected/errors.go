package protected

import (
	"fmt"
	"time"
)

type BDRAError struct {
	Code        string    `json:"code"`
	Message     string    `json:"message"`
	RingID      string    `json:"ringId"`
	Domain      string    `json:"domain"`
	Timestamp   time.Time `json:"timestamp"`
	Remediation string    `json:"remediation"`
}

// Error implements the standard Go error interface.
func (e *BDRAError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.Code, e.Domain, e.Message)
}

func NewWriteRejectedError(ringID, domain string) *BDRAError {
	return &BDRAError{
		Code:        "BDRA_WRITE_REJECTED",
		Message:     "Write operations are rejected while operating out of the fallback cache snapshot.",
		RingID:      ringID,
		Domain:      domain,
		Timestamp:   time.Now().UTC(),
		Remediation: "Do not queue transaction. Caller retains execution ownership and must retry post-incident.",
	}
}