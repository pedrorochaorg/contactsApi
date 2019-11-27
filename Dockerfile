# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.13-alpine as builder

# Maintainer information
LABEL Maintainer="Pedro Rocha <pedrorocha.org@gmail.com>"

# Project path
ENV PACKAGENAME 'contacts-api'
ENV PACKAGE github.com/pedrorochaorg/${PACKAGENAME}

# Environment Variables
ENV GOPATH /go
ENV GOPROJECT ${GOPATH}/src/${PACKAGE}

# Arguments Variables (custom)
ARG WATCH
ARG BUILDSTAMP='local'
ARG GITHASH='local'
ARG APP_VERSION='local'

# Add go files to the project
RUN mkdir -p ${GOPROJECT}
ADD . ${GOPROJECT}

# GO Project directory
WORKDIR ${GOPROJECT}

# Setup
RUN apk update && \
    apk --no-cache add ca-certificates

# Build application with custom ldflags
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod vendor -o /app ./cmd/webserver

ENTRYPOINT ["/app"]

# Build Scratch Image
FROM scratch

# Copy executable to scratch container
COPY --from=builder /app app

ENTRYPOINT ["./app"]