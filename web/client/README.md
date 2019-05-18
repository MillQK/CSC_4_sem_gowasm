### Client compilation
GOARCH=wasm GOOS=js go build -o web/prj.wasm client.go

### Page server
Run page server to serve index.html
