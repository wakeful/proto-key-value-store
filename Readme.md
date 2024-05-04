# storage service prototype.

this is a simple prototype of a key value (string -> string) storage service.
Build to practice Go stdlib and idiomatic code style, mainly interfaces.
Currently, only in-memory storage is implemented.

```shell
# start the service
go run cmd/main.go
```

## API interface

this service listens on port 8080 and provides `/storage` endpoint to store and retrieve key-value pairs (via POST & GET
http methods).

### let's store a number

```shell
curl -X POST --location "http://127.0.0.1:8080/storage" \
-d '{
"key": "number",
"value": "42"
}'
```

> {"key":"number","value":"42"}

### let's retrieve the number

```shell
curl -X GET --location "http://127.0.0.1:8080/storage?key=number"
```

> {"key":"number","value":"42"}

## Producing binary

you can use [GoReleaser](https://goreleaser.com/) to build the binary.

```shell
goreleaser build --clean --snapshot --clean
```
