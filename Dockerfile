FROM golang:1.25.6 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags "-w -extldflags -static" -o server .

FROM scratch 
COPY --from=builder /app/server /server
COPY --from=builder /app/internal /internal
ENTRYPOINT ["/server"]