## Config

Make sure to add the vercel.json file so the routing will work properly with vercel. This will rewrite all requests to the application to the `api/index.go` handler where the router will take over.

```bash
{
  "rewrites": [
    { "source": "(.*)", "destination": "api/index.go" }
  ]
}
```