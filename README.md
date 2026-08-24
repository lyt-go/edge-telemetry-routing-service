# Edge Telemetry Routing Service

This Go service accepts device telemetry, keeps short-lived snapshots, dispatches
alerts, coordinates collection jobs, and records firmware activation events. It
uses only the standard library so it can run in constrained edge environments.

## Layout

- `cmd/server`: HTTP health endpoint
- `internal/gateway`, `internal/transport`: request dispatch and cancellation
- `internal/ingest`, `internal/snapshot`, `internal/exporter`: batch ownership
- `internal/service`: business workflows
- `internal/collector`, `internal/alert`, `internal/firmware`: background jobs

## Run

```bash
go run ./cmd/server
```

The server listens on `:8080` by default. Set `ADDR` to use another address.

## Check

```bash
go build ./...
go test ./...
```
