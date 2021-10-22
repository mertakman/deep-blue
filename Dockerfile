FROM golang:1.17

RUN apt update 
RUN apt install git 
RUN apt install make 
RUN apt install openssl 

WORKDIR $GOPATH/src/github.com/mertakman/deep-blue

ENV GOBIN /usr/local/go/bin
ENV GOPROXY=direct
ENV GO111MODULE=on
ENV CGO_ENABLED=0

COPY ./ ./

RUN go mod download
RUN go test ./...

RUN go build -o deep-blue -i ./cmd/api/main.go

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /deep-blue/server

COPY --from=0 /go/src/github.com/mertakman/deep-blue/deep-blue ./

EXPOSE 8080

CMD ["./deep-blue"]
