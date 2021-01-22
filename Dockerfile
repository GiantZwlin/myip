FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build

FROM debian:buster-slim
WORKDIR /app
COPY --from=builder /app/myip .
EXPOSE 80
CMD ["./myip"]