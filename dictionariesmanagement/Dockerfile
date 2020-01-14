FROM golang:alpine as deps
RUN apk add --no-cache ca-certificates git
WORKDIR /app
COPY go.* /app/
RUN go mod download

FROM deps as builder
WORKDIR /app
COPY dictionariesmanagement/ /app/dictionariesmanagement/
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app dictionariesmanagement/main.go
FROM alpine
RUN apk add --no-cache ca-certificates
ENV GIN_MODE=release
COPY --from=builder /app/app /dictionariesmanagement
ENTRYPOINT ["/dictionariesmanagement"]
CMD ["--help"]
EXPOSE 8080
