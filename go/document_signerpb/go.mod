module github.com/paysuper/paysuper-proto/go/document_signerpb

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.8.0
	github.com/stretchr/testify v1.4.0
)

replace (
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
	github.com/micro/go-micro => github.com/micro/go-micro v1.8.0
)
