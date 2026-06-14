# Contributing to BDRA Lite

Thank you for your interest in contributing to the Business-Domain Ring Architecture (BDRA) Lite reference implementation! 

This repository enforces a strict, architecture-first code discipline. Unlike conventional codebases where code style is the primary concern, contributions here are evaluated primarily on **architectural boundary integrity**. 

Before submitting a Pull Request, please review our architectural guidelines and local development workflow below.

---

## 🏛️ The Core Rule: Layer Discipline

Every domain ring contains three strictly isolated layers. Your contribution **must** respect these boundaries:

1. **Pure (`*/pure/`)**: Zero-dependency domain logic. No I/O, no logging frameworks, no network calls, no `database/sql`, and no external context handling. If a function in this layer cannot be tested using raw inputs and outputs without mocks, it does not belong here.
2. **Protected (`*/protected/`)**: The contract boundary. Defines interfaces, cross-ring Data Transfer Objects (DTOs), and structural data types.
3. **Public (`*/public/`)**: The execution framework. This is where HTTP routing, gRPC servers, SQL queries, cache clients, and concrete infrastructure adapters live.

### Ring Dependency Rules
* Outer rings (e.g., Ring 2 Analytics) can import Protected interfaces from inner rings (e.g., Ring 1 Orders).
* **Inner rings can never import outer rings.**
* Horizontal imports inside the same ring are strictly limited: `protected` defines what `public` executes, and `public` feeds data into `pure`.

---

## 🛠️ Local Development Setup

### Prerequisites
* **Go**: Version 1.21 or higher
* **Docker & Docker Compose**: For baseline infrastructure dependencies (PostgreSQL/Redis)

### 1. Initialize the Environment
Clone the repository and spin up the supporting sandbox infrastructure:
```bash
docker-compose up -d
```

### 2. Running Tests
We lean heavily on Go's native testing toolchain. Because the `pure` layer is fully decoupled, unit tests are blisteringly fast and execute without any database dependencies.

**On Linux/macOS/Git Bash:**
```bash
make test
```

**On Windows (PowerShell):**
```powershell
go test -v ./...
```

### 3. Architecture Boundary Verification
Before pushing, verify that no banned infrastructure components have bled into your domain logic layers.

```powershell
# Run the local safety linter script
go run scripts/local_lint.go
```

---

## 📥 Code Contribution Workflow

We follow a typical fork-and-pull-request workflow:

1. **Fork the Repository**: Create your own copy of the codebase.
2. **Create a Feature Branch**: Use descriptive naming conventions (`feat/ring1-order-cancellation` or `fix/tier2-fallback-race-condition`).
3. **Commit with Intention**: Write clean, descriptive commit messages. Ensure all new logic in the `pure` layer has 100% test coverage.
4. **Run Local Checks**: Ensure `go fmt ./...` is run and that both the test suite and structural linter pass cleanly.
5. **Open a Pull Request**: Submit your PR against our `main` branch.

---

## 🔍 Pull Request Review Process

Every single Pull Request is subjected to both automated and manual gate checks:

* **Automated CI Pipeline**: GitHub Actions will automatically compile the codebase, run all unit tests, and execute the structural boundary checker (`bdracheck`). If a boundary leak is identified, the build will immediately fail.
* **Manual Architecture Review**: Maintainers will look closely at your import blocks. PRs that introduce bypasses (such as adding global state, importing a database client directly into a core entity, or violating inward ring dependencies) will be rejected with feedback on how to refactor into the proper layer.

### Architectural Cleanliness Check-Off
When creating your PR, you will be prompted to fill out our compliance checklist. Be prepared to verify that your `pure` layers remain entirely deterministic and decoupled from infrastructure side-effects.

---

## 💬 Getting Help

If you are unsure whether a specific feature belongs in a `pure` package or a `public` structural adapter, open a GitHub Issue with the tag `architecture-discussion`. We prefer discussing the structural boundaries up front rather than requesting massive refactors during a code review!