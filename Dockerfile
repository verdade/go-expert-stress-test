FROM golang:1.22 as build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/main cmd/main.go

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch


WORKDIR /app

COPY --from=build /app/bin/main .

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt


ENTRYPOINT ["./main"]