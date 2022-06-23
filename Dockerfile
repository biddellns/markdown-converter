FROM golang:1.18 as builder

WORKDIR /go/src/interview-markdown-converter
COPY cmd ./cmd/
COPY lib ./lib/
COPY go.mod go.sum ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cli ./cmd/cli/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/interview-markdown-converter/cli .
ENTRYPOINT ["./cli"]