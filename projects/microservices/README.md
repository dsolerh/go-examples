# microservices

Educational microservices app in Go. Each service lives in `services/<name>`
with its own `go.mod`. Each service is listed in the repo root `go.work`.

## Services

| Service | Port | Purpose |
|---------|------|---------|
| auth    | 8081 | Register users, log them in, issue and validate JWTs |

## Auth service

Run it:

```sh
cd services/auth
go run ./cmd/auth
```

Env vars: `AUTH_ADDR` (`:8081`), `AUTH_JWT_SECRET`, `AUTH_ACCESS_TTL` (`15m`), `AUTH_REFRESH_TTL` (`168h`).

Try it:

```sh
curl -X POST localhost:8081/register -d '{"email":"ana@example.com","password":"supersecret"}'
curl -X POST localhost:8081/login    -d '{"email":"ana@example.com","password":"supersecret"}'
curl localhost:8081/validate -H "Authorization: Bearer <access_token>"
curl -X POST localhost:8081/refresh  -d '{"refresh_token":"<refresh_token>"}'
```

Layout:

```
cmd/auth/          main: wiring + graceful shutdown
internal/config/   env config
internal/user/     user model + Repository interface (in-memory for now)
internal/token/    JWT issue/parse (HS256)
internal/service/  business logic, no HTTP
internal/httpapi/  HTTP handlers, error -> status mapping
```
