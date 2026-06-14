@echo off
title BDRA Lite Symmetric Workspace Initializer
echo =======================================================
echo   Initializing Symmetrically Complete BDRA Lite Layout
echo =======================================================

:: Create All Structural Sub-Packages
echo [1/3] Generating isolated directory ring structures...
mkdir .github\workflows
mkdir cmd\app
mkdir internal\ring0\pure
mkdir internal\ring0\protected
mkdir internal\ring0\public
mkdir internal\ring1\pure
mkdir internal\ring1\protected
mkdir internal\ring1\public
mkdir internal\ring1\health
mkdir internal\ring2\pure
mkdir internal\ring2\protected
mkdir internal\ring2\public
mkdir internal\ring2\health

:: Seed Code Placement Tokens
echo [2/3] Placing code modules and configurations...
copy NUL cmd\app\main.go > NUL
copy NUL internal\ring0\pure\user.go > NUL
copy NUL internal\ring0\protected\ports.go > NUL
copy NUL internal\ring0\public\service.go > NUL
copy NUL internal\ring1\pure\order.go > NUL
copy NUL internal\ring1\pure\order_test.go > NUL
copy NUL internal\ring1\protected\errors.go > NUL
copy NUL internal\ring1\protected\ports.go > NUL
copy NUL internal\ring1\public\adapters.go > NUL
copy NUL internal\ring1\public\http.go > NUL
copy NUL internal\ring1\public\postgres.go > NUL
copy NUL internal\ring1\public\service.go > NUL
copy NUL internal\ring1\health\health.go > NUL
copy NUL internal\ring1\health\supervisor.go > NUL
copy NUL internal\ring2\pure\analytics.go > NUL
copy NUL internal\ring2\protected\ports.go > NUL
copy NUL internal\ring2\public\service.go > NUL
copy NUL internal\ring2\health\health.go > NUL
copy NUL .github\workflows\bdra-lint.yml > NUL
copy NUL Makefile > NUL

:: Build Local Configuration Primitives
(
echo {
echo   "$schema": "https://raw.githubusercontent.com/bdra-io/bdra-spec/main/schemas/v1/linter.json",
echo   "projectName": "modular-monolith",
echo   "architectureStrategy": "BDRA-Lite",
echo   "rules": {
echo     "disallowExternalImportsInPure": {
echo       "targetDirs": ["internal/ring*/pure/**/*.go"],
echo       "allowedPrefixes": ["modular-monolith/internal/ring*/pure"]
echo     },
echo     "disallowIOInProtected": {
echo       "targetDirs": ["internal/ring*/protected/**/*.go"],
echo       "forbiddenPackages": ["database/sql", "net/http", "os", "io"]
echo     },
echo     "enforceInwardDependencyFlow": {
echo       "rings": [
echo         { "id": "ring0", "path": "internal/ring0" },
echo         { "id": "ring1", "path": "internal/ring1", "allowedDependencies": ["ring0"] },
echo         { "id": "ring2", "path": "internal/ring2", "allowedDependencies": ["ring0", "ring1"] }
echo       ]
echo     }
echo   }
echo }
) > bdracheck.json

(
echo module github.com/bdra-io/monolith
echo.
echo go 1.22
echo.
echo require github.com/lib/pq v1.10.9
) > go.mod

(
echo version: '3.8'
echo.
echo services:
echo   postgres:
echo     image: postgres:15-alpine
echo     environment:
echo       POSTGRES_USER: user
echo       POSTGRES_PASSWORD: pass
echo       POSTGRES_DB: db
echo     ports:
echo       - "5432:5432"
echo     volumes:
echo       - pgdata:/var/lib/postgresql/data
echo.
echo   redis:
echo     image: redis:7-alpine
echo     ports:
echo       - "6379:6379"
echo.
echo volumes:
echo   pgdata:
) > docker-compose.yml

echo [3/3] Environmental initialization successfully completed!
echo =======================================================
pause