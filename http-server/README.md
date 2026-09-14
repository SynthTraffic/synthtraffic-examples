# HTTP test server

A small Go server that accepts the requests produced by [`../connectors/http.yaml`](../connectors/http.yaml). It listens on `localhost:18081`, returns JSON `202 Accepted`, and logs method, path, and status (not headers or query values).

Start it in one terminal:

```bash
cd http-server
go run .
```

Then run Synthtraffic in another:

```bash
export API_TOKEN=test-key
synthtraffic run connectors/http.yaml
```

PowerShell:

```powershell
$env:API_TOKEN = "test-key"
.\synthtraffic.exe run connectors\http.yaml
```

Preview without the server:

```bash
synthtraffic run connectors/http.yaml --stdout --events 3 --seed 42
```

Endpoints:

- `GET /health` returns `200 OK`
- `/api/orders/{orderId}` accepts an optional JSON body
- `/status/404` and `/status/500` return sample error statuses. Synthtraffic still treats these as completed requests — the HTTP connector ignores response status and content.

Set `HTTP_TEST_SERVER_ADDRESS` to override `:18081`, and point the scenario `baseUrl` at the same port.
