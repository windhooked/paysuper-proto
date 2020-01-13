protoc -I=../declarations/billing --micro_out=./billingpb/ --go_out=./billingpb/ ../declarations/billing/billing.proto
protoc -I=../declarations/billing --micro_out=./billingpb/ --go_out=./billingpb/ ../declarations/billing/cardpay.proto
protoc -I=../declarations/billing --micro_out=./billingpb/ --go_out=./billingpb/ ../declarations/billing/grpc.proto
protoc -I=../declarations/billing --micro_out=./billingpb/ --go_out=./billingpb/ ../declarations/billing/paylink.proto
