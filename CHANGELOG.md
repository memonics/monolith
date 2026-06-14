# Changelog

All notable changes to the BDRA Lite reference implementation will be documented in this file. This project adheres strictly to Semantic Versioning (`MAJOR.MINOR.PATCH`).

---

## [1.5.1] — 2026-06-13

### Fixed
* **Concurrency Protection:** Migrated the service coordinator's internal degradation state fields from standard booleans to `sync/atomic.Bool` to eliminate multi-threaded memory data races under concurrent HTTP loads.
* **Error Verification Engine:** Explicitly matched the custom `BDRAError` structural block to satisfy Go's native `error` interface, resolving a bug where downstream `errors.As` unpackers silently dropped domain anomaly signatures.
* **Resource Leaks:** Wrapped background out-of-band monitoring execution loops in proper context cancellation logic with clean `ticker.Stop()` deferrals to prevent heap accumulation.

### Added
* **Cache Warming Primitives:** Added transparent, push-based local snapshot serialization directly into successful primary repo execution paths.
* **Symmetric Architecture Expansion:** Implemented Ring 2 (`Operational Intelligence`) along with its own pure logic packages and independent out-of-band health servers to fully manifest multi-ring integration rules.

---

## [1.5.0] — 2026-06-01
* Initial open-source release of the BDRA Lite modular monolith blueprint specification tier.