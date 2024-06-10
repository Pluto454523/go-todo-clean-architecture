FROM golang:1.22.0-alpine AS builder
WORKDIR $GOPATH/src/app/
ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /go/bin/app cmd/generics_server/*.go

FROM alpine
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app/app
ENTRYPOINT ["/app/app"]
