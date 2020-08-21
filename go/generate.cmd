protoc -I=../declarations/casbin -I=%GOPATH%/src/github.com/envoyproxy/ -I=%GOPATH%/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.7/ -I=%GOPATH%/src/github.com/googleapis/googleapis/ --micro_out=./casbinpb/ --go_out=./casbinpb/ ../declarations/casbin/casbin.proto

protoc -I=../declarations/billing --micro_out=./billingpb/ --go_out=./billingpb/ ../declarations/billing/billing.proto
protoc -I=../declarations/billing --micro_out=./billingpb/ --go_out=./billingpb/ ../declarations/billing/cardpay.proto
protoc -I=../declarations/billing --micro_out=./billingpb/ --go_out=./billingpb/ ../declarations/billing/grpc.proto
protoc -I=../declarations/billing --micro_out=./billingpb/ --go_out=./billingpb/ ../declarations/billing/paylink.proto

protoc -I=../declarations/currencies-service --micro_out=./currenciespb/ --go_out=./currenciespb/ ../declarations/currencies-service/currencies.proto

protoc -I=../declarations/recurring-service --micro_out=./recurringpb/ --go_out=./recurringpb/ ../declarations/recurring-service/entity.proto
protoc -I=../declarations/recurring-service --micro_out=./recurringpb/ --go_out=./recurringpb/ ../declarations/recurring-service/repository.proto
protoc -I=../declarations/recurring-service --micro_out=./recurringpb/ --go_out=./recurringpb/ ../declarations/recurring-service/xsolla.proto

protoc -I=../declarations/postmark-sender --micro_out=./postmarkpb/ --go_out=./postmarkpb/ ../declarations/postmark-sender/postmark.proto

protoc -I=../declarations/reporter --micro_out=./reporterpb/ --go_out=./reporterpb/ ../declarations/reporter/reporter.proto

protoc -I=../declarations/tax-service --micro_out=./taxpb/ --go_out=./taxpb/ ../declarations/tax-service/tax_service.proto

protoc -I=../declarations/document-signer --micro_out=./document_signerpb/ --go_out=./document_signerpb/ ../declarations/document-signer/signer.proto

protoc -I=../declarations/notifier --micro_out=./notifierpb/ --go_out=./notifierpb/ ../declarations/notifier/notifier.proto

