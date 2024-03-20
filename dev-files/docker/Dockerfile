FROM golang:1.22 as builder
LABEL maintainer="az <azusachino@duck.com>"
WORKDIR /app
COPY go.mod go.sum ./
# for cn setup proxy
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct
RUN ["go", "mod", "download"]

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ficus .

# setup running environment
FROM alpine:latest
# replace mirror at first
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add --update --no-cache ca-certificates
WORKDIR /root
COPY --from=builder /app/ficus .
EXPOSE 8090
ENTRYPOINT ["./ficus"]
