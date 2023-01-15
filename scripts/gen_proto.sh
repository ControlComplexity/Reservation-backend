protoc --go_out=. ./proto/reservation.proto
protoc --go-grpc_out=. ./proto/reservation.proto
protoc --grpc-gateway_opt=logtostderr=true --grpc-gateway_out=. ./proto/reservation.proto
protoc --openapiv2_opt=logtostderr=true,json_names_for_fields=false --openapiv2_out=./ ./proto/reservation.proto
protoc  --validate_out="lang=go:./"  ./proto/reservation.proto
