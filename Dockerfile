FROM golang:1.25.6 as builder
WORKDIR /app

ARG VERSION
ARG COMMIT
ARG BUILD_TIME

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo \
    -ldflags "\
        -w -extldflags -static \
        -X 'api-auth/internal/version.Version=${VERSION}' \
        -X 'api-auth/internal/version.Commit=${COMMIT}' \
        -X 'api-auth/internal/version.BuildTime=${BUILD_TIME}'" \
    -o server .

FROM scratch 
COPY --from=builder /app/server /server
COPY --from=builder /app/internal /internal
ENTRYPOINT ["/server"]