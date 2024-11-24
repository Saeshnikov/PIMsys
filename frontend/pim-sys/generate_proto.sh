cd src/grpc/sso
protoc --js_out=import_style=commonjs,binary:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. sso.proto
cd ../../../

cd src/grpc/shop
protoc --js_out=import_style=commonjs,binary:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. shop.proto