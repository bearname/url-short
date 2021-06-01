FROM golang:1.15.6 AS builder
WORKDIR /go/src/url-short
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /go/src/url-short/cmd/short
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/src/url-short/bin/short

FROM alpine:3.12.3
WORKDIR /app
COPY --from=builder /go/src/url-short/bin/short /app/short

EXPOSE 8080
ENTRYPOINT ["/app/short"]