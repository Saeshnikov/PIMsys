FROM ubuntu:22.04

RUN apt update && apt upgrade
RUN apt install -y wget make
RUN wget https://go.dev/dl/go1.23.3.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.23.3.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

RUN apt install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN PATH="$PATH:$(go env GOPATH)/bin"
