FROM golang:1.19
WORKDIR /go/src
COPY . .
RUN go get -d -v ./...
RUN go build
EXPOSE 8082
ENTRYPOINT ["./image-service"]
