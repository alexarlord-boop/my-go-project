# Client Generation

We can generate client code for the API. The following command generates a client for the API:

```bash
# from the client directory
 swagger generate client -f ../product-api/swagger.yaml -A product-api
```

Also, I had to get this package to resolve generation errors:

```bash
 go get github.com/go-openapi/runtime/client@v0.28.0
```
