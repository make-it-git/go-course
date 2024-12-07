FROM golang:1.23 as builder
WORKDIR /app
COPY go.mod ./
COPY main.go .
RUN go build -o app .

FROM debian:bullseye
RUN apt-get update && apt-get install -y \
    strace \
    bash \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/app .
CMD ["bash"]

