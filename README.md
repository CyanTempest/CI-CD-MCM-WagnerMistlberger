![CI Status](https://github.com/CyanTempest/CI-CD-MCM-WagnerMistlberger/actions/workflows/ci.yml/badge.svg)

# Exercise 2: Microservice Architecture, Docker & GitHub Actions

**Course:** Continuous Delivery in Agile Software Development (Master)
**Points:** 24

## Learning Objectives

- Extend a CI pipeline with quality gates and code analysis
- Configure SonarCloud for static code analysis and coverage tracking
- Use matrix builds to test across multiple Go versions
- Integrate linting with golangci-lint
- Understand code quality metrics and technical debt

## Prerequisites

- Completed Exercise 1
- Docker Desktop installed
- Basic understanding of REST APIs

## Project Overview

The Product Catalog API has been extended with:

- **PostgreSQL storage** (`internal/store/postgres.go`) -- persistent database backend
- **Dockerfile** -- multi-stage build for minimal container image
- **docker-compose.yml** -- orchestrates API + PostgreSQL
- **GitHub Actions** (`.github/workflows/ci.yml`) -- basic CI pipeline

### Architecture

```
┌──────────────┐     ┌──────────────┐
│   Client     │────▶│   API (Go)   │
│  (curl/HTTP) │     │   Port 8080  │
└──────────────┘     └──────┬───────┘
                            │
                     ┌──────▼───────┐
                     │  PostgreSQL  │
                     │  Port 5432   │
                     └──────────────┘
```

### Local Development

```bash
# Run with in-memory store (no Docker needed)
go run ./cmd/api

# Run with Docker Compose (API + PostgreSQL)
docker compose up --build

# Test the API
curl http://localhost:8080/health
curl http://localhost:8080/products
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Widget","price":9.99}'
```

---

## Tasks

### Task 1: Understand the Architecture (2 Points)

**Expected result:** 4 parallel test jobs (2 Go versions x 2 OS).

**Deliverable:** Screenshot of the GitHub Actions matrix view showing all jobs.

---

### Task 2: Complete the GitHub Actions Workflow (6 Points)

1. **Add a `lint` job** to the CI workflow that:
   - Runs `golangci-lint` using the `golangci/golangci-lint-action@v4` action
   - Uses the `.golangci.yml` configuration file
   - Runs in parallel with the test matrix (does not depend on `test`)

2. **Enable additional linters** in `.golangci.yml` (see TODOs):
   - `gofmt` -- enforces standard Go formatting
   - `gocyclo` -- detects overly complex functions
   - `misspell` -- catches common typos
   - `gocritic` -- advanced Go code analysis

3. **Fix any linting issues** that are reported in the existing code.

**Deliverable:** Clean lint run (no warnings). Screenshot of the lint job passing.

---

### Task 3: Docker & Docker Compose (8 Points)

1. **Analyze the Dockerfile:**

   - Explain each stage of the multi-stage build. Why two stages?
   - What does `CGO_ENABLED=0` do and why is it important?
   - What is the final image size? Compare it to a single-stage build.

2. **Run the application with Docker Compose:**

   ```bash
   docker compose up --build
   ```

3. **Test all CRUD operations** using `curl` or a tool like Postman:

   - Create at least 3 products
   - List all products
   - Update a product
   - Delete a product
   - Verify the product is gone

4. **Verify data persistence:**

   - Stop and restart the containers (`docker compose down` then `up`)
   - Check if the products still exist (they should, thanks to the volume)

4. **Add the `SONAR_TOKEN` secret** to your repository settings.

5. **Review the SonarCloud dashboard:**
   - What is the code coverage percentage?
   - Are there any code smells or bugs detected?
   - What is the technical debt estimate?

**Deliverable:** Link to your SonarCloud project dashboard. Screenshot showing the quality gate result.

---

### Task 4: Code Coverage Improvement (6 Points)

1. **Check current coverage:**
   ```bash
   go test -coverprofile=coverage.out ./...
   go tool cover -func=coverage.out
   go tool cover -html=coverage.out -o coverage.html
   ```

2. **Improve coverage to at least 80%** by adding tests for uncovered code paths. Focus on:
   - Edge cases in handlers (invalid IDs, malformed JSON)
   - Error paths in the store layer
   - The `Validate()` method edge cases

3. **Add a coverage threshold check** to the CI pipeline as a step after running tests:
   - Extract the total coverage percentage from `go tool cover -func`
   - Fail the build if coverage is below 80%
   - Use `::error::` to display the error in the GitHub Actions UI

   > **Hint:** `go tool cover -func=coverage.out | grep total` gives you the total line. Use `awk` and `sed` to extract the number. Use `bc` for the comparison (works on both Linux and macOS).

| Method | Endpoint         | Description       | Request Body                  |
| ------ | ---------------- | ----------------- | ----------------------------- |
| GET    | `/health`        | Health check      | --                            |
| GET    | `/products`      | List all products | --                            |
| POST   | `/products`      | Create product    | `{"name":"...","price":0.00}` |
| GET    | `/products/{id}` | Get product by ID | --                            |
| PUT    | `/products/{id}` | Update product    | `{"name":"...","price":0.00}` |
| DELETE | `/products/{id}` | Delete product    | --                            |

---

## Grading

| Task                       | Points |
| -------------------------- | ------ |
| Architecture Documentation | 2      |
| GitHub Actions Workflow    | 6      |
| Docker & Docker Compose    | 8      |
| Handler Tests              | 8      |
| **Total**                  | **24** |

## Author

- FH-Prof. Dr. Marc Kurz (marc.kurz@fh-hagenberg.at)

