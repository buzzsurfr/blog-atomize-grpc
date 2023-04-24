module github.com/buzzsurfr/blog-atomize-grpc

go 1.20

replace github.com/buzzsurfr/blog-atomize-grpc/loyalty => ./loyalty
replace github.com/buzzsurfr/blog-atomize-grpc/orders => ./orders

require (
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
)
