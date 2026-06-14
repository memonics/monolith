#!/bin/bash
set -e

echo "=== BDRA Lite Resilience Sandbox Simulator ==="
echo "1. Triggering database infrastructure fault (Tier 2/3)..."
curl -X POST http://localhost:8080/simulate/fault
echo -e "\n\n--> Attempting a write operation on /orders while degraded:"
curl -i -X POST -H "Content-Type: application/json" -d '{"user_id":"usr_007","amount":250.00}' http://localhost:8080/orders

echo -e "\n\n2. Recovering baseline operations (Tier 1)..."
curl -X POST http://localhost:8080/simulate/recover
echo -e "\n\n=== Simulation Sequence Complete ==="