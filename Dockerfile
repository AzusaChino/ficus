FROM golang:latest as builder
LABEL maintainer="az <azusachino@yahoo.com>"
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
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /app/ficus .
EXPOSE 8090
ENTRYPOINT ["./ficus"]