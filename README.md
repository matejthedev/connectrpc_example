# Connect-Go example

A minimal [Connect](https://connectrpc.com/) RPC service in Go: one `GreetService` with two unary methods, served over plain HTTP/2 (h2c) so it works with gRPC, gRPC-Web and plain JSON clients from the same handler.

The point of the repo is the wiring, not the service: Buf-managed code generation, the Connect handler mounted on a standard `net/http` mux, and h2c so no TLS is needed locally.

## Layout

```
greet/v1/greet.proto        service definition
buf.yaml, buf.gen.yaml      Buf module + codegen config (protoc-gen-go, protoc-gen-connect-go)
gen/greet/v1/               generated message and Connect code (committed)
cmd/server/main.go          the server
install.sh                  installs buf, grpcurl and the two protoc plugins
shell.nix                   Go toolchain for Nix users
```

## Run

```sh
go run ./cmd/server          # listens on localhost:8888
```

Call it with plain JSON, no client code needed:

```sh
curl -X POST http://localhost:8888/greet.v1.GreetService/Greet \
  -H 'Content-Type: application/json' -d '{"name": "World"}'
# {"greeting":"Hello, World!"}

curl -X POST http://localhost:8888/greet.v1.GreetService/Mul \
  -H 'Content-Type: application/json' -d '{"inputNum": 4}'
# {"result":40}
```

Or with gRPC, via grpcurl (h2c, so `-plaintext`):

```sh
grpcurl -plaintext -proto greet/v1/greet.proto -d '{"name": "World"}' \
  localhost:8888 greet.v1.GreetService/Greet
```

## Regenerate code

```sh
./install.sh   # once
buf generate
```
