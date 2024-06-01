FROM golang:1.22.3-alpine3.20 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLE=0 GOOS=linux go build -o myapp

FROM alpine:3.20

COPY --from=builder /app/myapp /

EXPOSE 8000

ENTRYPOINT ["/myapp"]
