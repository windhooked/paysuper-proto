module github.com/paysuper/paysuper-proto/go/casbinpb

go 1.13

require (
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/golang/protobuf v1.3.2
	github.com/grpc-ecosystem/grpc-gateway v1.12.1
	github.com/micro/go-micro v1.18.0
	github.com/stretchr/testify v1.4.0
	google.golang.org/genproto v0.0.0-20200117163144-32f20d992d24
)

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
