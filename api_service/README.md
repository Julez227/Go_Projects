# api_service

A dependency-free Go HTTP microservice starter. No business logic — it's the
skeleton (routing, middleware, config, health checks) meant to be extended.

## Run it

```
API_KEY=$(openssl rand -hex 32) go run ./api_service
```

`API_KEY` is required; the service refuses to start without one (fail closed,
no baked-in default credential).

## Endpoints

| Route         | Auth | Purpose                                    |
|---------------|------|---------------------------------------------|
| `GET /healthz`| none | liveness probe                              |
| `GET /readyz` | none | readiness probe (wire up dependency checks) |
| `GET /v1/example` | API key | example protected route to copy for real handlers |

Send the key as `X-API-Key: <key>` or `Authorization: Bearer <key>`.

## What's already handled

- **Panic recovery** — a panic anywhere in a handler returns a generic 500
  instead of crashing the process or leaking a stack trace.
- **Auth** — constant-time API key comparison (`crypto/subtle`) to avoid
  timing side channels. Swap for JWT/OAuth by replacing `requireAPIKey` in
  [middleware.go](middleware.go) — use a vetted library (e.g.
  `github.com/golang-jwt/jwt`) rather than hand-rolling token verification.
- **Rate limiting** — per-client-IP token bucket (`RATE_LIMIT_RPS` /
  `RATE_LIMIT_BURST`), with a background sweep of stale buckets so it can't
  become an unbounded-memory DoS vector itself.
- **Request size limits** — `MAX_REQUEST_BYTES` (default 1 MiB) via
  `http.MaxBytesReader`.
- **Timeouts everywhere** — read, read-header, write, idle, and a whole-request
  timeout (`http.TimeoutHandler`), plus graceful shutdown on
  `SIGINT`/`SIGTERM`.
- **Security headers** — `X-Content-Type-Options`, `X-Frame-Options`,
  `Content-Security-Policy: default-src 'none'`, `Strict-Transport-Security`,
  `Cache-Control: no-store`.
- **CORS** — deny by default; only origins listed in `ALLOWED_ORIGINS` are
  permitted.
- **Structured logging** — JSON logs via `log/slog`, one line per request with
  a per-request ID (`X-Request-ID`) for tracing across services.

## Config (environment variables)

| Var | Default | Notes |
|---|---|---|
| `PORT` | `8080` | |
| `API_KEY` | *(required)* | min 16 chars |
| `ALLOWED_ORIGINS` | *(none)* | comma-separated |
| `MAX_REQUEST_BYTES` | `1048576` | |
| `RATE_LIMIT_RPS` / `RATE_LIMIT_BURST` | `5` / `10` | per client IP |
| `REQUEST_TIMEOUT` | `10s` | whole-request deadline |
| `READ_TIMEOUT` / `READ_HEADER_TIMEOUT` / `WRITE_TIMEOUT` / `IDLE_TIMEOUT` | `5s`/`5s`/`10s`/`60s` | |
| `SHUTDOWN_TIMEOUT` | `15s` | grace period for in-flight requests on shutdown |

## Deployment notes

- **TLS**: this service listens on plain HTTP (`ListenAndServe`) by default,
  expecting TLS to be terminated by an upstream load balancer, ingress, or
  service-mesh sidecar (typical in a microservices deployment). If you need
  the service to terminate TLS itself, switch to
  `srv.ListenAndServeTLS(certFile, keyFile)` — `TLSConfig.MinVersion` is
  already pinned to TLS 1.2.
- **Secrets**: `API_KEY` should come from a secrets manager / orchestrator
  secret, not a checked-in `.env` file.
- **Dependency scanning**: run `go vet ./...` and
  [`govulncheck`](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck)
  regularly, especially once real dependencies get added.
- **Rate limiting is per-instance**: the token bucket lives in process memory,
  so it resets on restart and doesn't coordinate across replicas. For
  multi-replica deployments, front the service with a shared limiter (API
  gateway, Redis-backed limiter) if a tighter global bound is needed.
