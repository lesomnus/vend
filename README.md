# vend

Go module vanity import server. Responds to `?go-get=1` requests with the appropriate `go-import` meta tag so `go get` can resolve custom import paths to their VCS roots.

## Configuration

Place a `vend.yaml` (or `vend.yml`) in the working directory, or pass `--config <path>`.

```yaml
server:
  addr: ":8080"

packages:
  x/net:
    target: https://github.com/golang/net
    branch: master
  x/sync:
    target: https://github.com/golang/sync
    branch: master
    hidden: true   # suppresses pkg.go.dev indexing
```

### `packages` fields

| Field    | Type   | Description                                             |
| -------- | ------ | ------------------------------------------------------- |
| `target` | string | VCS repository root URL                                 |
| `branch` | string | Default branch (used for source links)                  |
| `hidden` | bool   | If true, adds `<meta name="robots" content="noindex"/>` |

## Endpoints

| Path                   | Description                                        |
| ---------------------- | -------------------------------------------------- |
| `GET /`                | Health check — always returns `200 OK`, no logging |
| `GET /<path>?go-get=1` | Vanity import response for the given package path  |

Any request without `?go-get=1`, or for an unknown package, returns `404`.
