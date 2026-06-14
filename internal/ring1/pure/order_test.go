package pure_test

import (
	"testing"
	"github.com/bdra-io/monolith/internal/ring1/pure"
)

func TestCreateOrder_ValidInvariants(t *testing.T) {
	order, err := pure.CreateOrder("ord_999", "usr_007", 1500.25)
	
	if err != nil {
		t.Fatalf("Expected clean domain creation, got structural error: %v", err)
	}
	if order.Status != "PENDING" {
		t.Errorf("Expected status PENDING, caught: %s", order.Status)
	}
	if order.Amount != 1500.25 {
		t.Errorf("Expected amount 1500.25, caught: %f", order.Amount)
	}
}

func TestCreateOrder_RejectsInvalidAmount(t *testing.T) {
	_, err := pure.CreateOrder("ord_999", "usr_007", -50.00)
	
	if err == nil {
		t.Fatal("Expected validation boundary error, but initialization succeeded blindly")
	}
	if err != pure.ErrInvalidAmount {
		t.Errorf("Expected ErrInvalidAmount, caught alternative payload: %v", err)
	}
}