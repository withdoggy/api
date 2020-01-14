FROM golang:alpine as deps
ENV GO111MODULE=on
RUN apk add --no-cache ca-certificates git ca-certificates tzdata && update-ca-certificates
RUN adduser -D -g '' withdoggy
WORKDIR /app
COPY go.* /app/
RUN go mod download && go mod verify

FROM deps as builder
WORKDIR /app
COPY dictionariesmanagement/ /app/dictionariesmanagement/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -installsuffix cgo -o app dictionariesmanagement/main.go

FROM scratch
ENV GIN_MODE=release
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/app /dictionariesmanagement
USER withdoggy
ENTRYPOINT ["/dictionariesmanagement"]
CMD ["--help"]
EXPOSE 8080
