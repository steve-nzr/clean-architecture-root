# Go version & Linux distribution
ARG GO_VERSION=1.18.2
ARG LINUX_DISTRO=bullseye

## Dependencies stage
FROM golang:${GO_VERSION}-${LINUX_DISTRO} AS dependencies
WORKDIR /usr/src/app

# Copy only module files for caching
COPY go.mod go.sum ./

# Pull dependencies
RUN go mod download

## Build stage
FROM dependencies as builder

# Copy rest of source code
COPY . .

# Build the app
RUN CGO_ENABLED=0 go build -o app main.go

## Runtime stage using distroless static, which contains ca-certificates, tzdata..
FROM gcr.io/distroless/static:nonroot as runtime

COPY --from=builder /usr/src/app/app /

CMD [ "/app", "serve" ]
