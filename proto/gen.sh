protoc --go_out=. ./proto/reservation.proto
protoc --go-grpc_out=. ./proto/reservation.proto
protoc --grpc-gateway_out=. ./proto/reservation.proto
protoc --openapiv2_opt=logtostderr=true --openapiv2_out=. ./proto/reservation.proto
protoc  --validate_out="lang=go:../generated"  ./proto/reservation.proto
