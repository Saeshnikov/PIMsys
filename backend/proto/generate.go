package generate

//go:generate protoc ./shop/shop.proto --go_out=../gen/go/ --go_opt=paths=source_relative --go-grpc_out=../gen/go/ --go-grpc_opt=paths=source_relative
//go:generate protoc ./sso/sso.proto --go_out=../gen/go/ --go_opt=paths=source_relative --go-grpc_out=../gen/go/ --go-grpc_opt=paths=source_relative
//go:generate protoc ./branch/branch.proto --go_out=../gen/go/ --go_opt=paths=source_relative --go-grpc_out=../gen/go/ --go-grpc_opt=paths=source_relative
//go:generate protoc ./products/products.proto --go_out=../gen/go/ --go_opt=paths=source_relative --go-grpc_out=../gen/go/ --go-grpc_opt=paths=source_relative
//go:generate protoc ./logs/logs.proto --go_out=../gen/go/ --go_opt=paths=source_relative --go-grpc_out=../gen/go/ --go-grpc_opt=paths=source_relative
