# Stage 1
FROM golang:1.22.4-alpine3.20 AS BUILD
# FROM golang:alpine AS BUILD

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /src

# Cache dependencies
COPY go.mod .
# COPY go.sum .
RUN go mod download

# Copy actual source
COPY . .

# Optional vulncheck
# RUN go install -v golang.org/x/vuln/cmd/govulncheck@latest

RUN CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags='-s -w -extldflags "-static"' -o ./out/ ./...

# Stage 2
FROM alpine:latest

WORKDIR /app

RUN apk update \
  && apk -U upgrade \
  && apk add --no-cache ca-certificates \
  && update-ca-certificates --fresh \
  && rm -rf /var/cache/apk/*

RUN addgroup gopher_group && adduser -S gopher -u 1000 -G gopher_group

COPY --from=BUILD /src/out/api cma

RUN chmod +x ./cma

USER gopher

ENV PORT=8080
EXPOSE ${PORT}

ENTRYPOINT ["/app/cma"]
