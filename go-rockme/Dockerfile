FROM golang:alpine AS builder

WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
RUN go build -o /go/app

FROM alpine:3.14
RUN mkdir -p /app
WORKDIR /go
COPY --from=builder /go/app /go
CMD ["./app"]