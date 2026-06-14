Write-Host "=== BDRA Lite Resilience Sandbox Simulator ===" -ForegroundColor Cyan

Write-Host "1. Triggering database infrastructure fault (Tier 2/3)..." -ForegroundColor Yellow
curl.exe -X POST http://localhost:8080/simulate/fault

Write-Host "`n`n--> Attempting a write operation on /orders while degraded:" -ForegroundColor Red
curl.exe -i -X POST -H "Content-Type: application/json" -d '{"user_id":"usr_007","amount":250.00}' http://localhost:8080/orders

Write-Host "`n`n2. Recovering baseline operations (Tier 1)..." -ForegroundColor Green
curl.exe -X POST http://localhost:8080/simulate/recover

Write-Host "`n`n=== Simulation Sequence Complete ===" -ForegroundColor Cyan