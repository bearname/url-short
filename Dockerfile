FROM golang:1.16.2 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o ./bin/urlshort ./cmd/short

######## Start a new stage #######
FROM alpine:3.11.5
RUN adduser -D app
USER app

COPY --from=builder /app/bin/urlshort /app/bin/

WORKDIR /app/

EXPOSE 8000
CMD ["./bin/urlshort"]