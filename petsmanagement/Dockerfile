FROM golang:alpine as deps
RUN apk add --no-cache ca-certificates git
# Copy local code to the container image.
WORKDIR /app
COPY go.* /app/
RUN go mod download

FROM deps as builder

# Copy local code to the container image.
WORKDIR /app
COPY petsmanagement/ /app/petsmanagement/

# Build the command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app petsmanagement/main.go

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine
RUN apk add --no-cache ca-certificates
ENV GIN_MODE=release
# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/app /petsmanagement

# Run the web service on container startup.
ENTRYPOINT ["/petsmanagement"]
CMD ["--help"]
EXPOSE 8080
