FROM golang:1.22 AS builder

ENV CGO_ENABLE 0 \
    GOPROXY https://goproxy.cn,direct \
    GO111MODULE=on

WORKDIR /biuld

COPY . .

RUN go build -ldflags="-s -w" -o /app/main main.go

FROM scratch
WORKDIR /app

COPY --from=builder /app/main /app/main

CMD ["./main"]