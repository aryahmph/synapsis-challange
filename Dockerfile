FROM golang:1.20-alpine3.18 as builder
WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build -o bin/web ./cmd/web

EXPOSE 4001

FROM alpine:3.18
COPY --from=builder /build/bin .
CMD ["./web"]