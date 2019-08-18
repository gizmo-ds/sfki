
#build stage
FROM golang:alpine AS builder
RUN mkdir -p /home/sfki
WORKDIR /home/sfki
COPY . /home/sfki
# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache git
RUN go build

#final stage
FROM alpine:latest
RUN mkdir -p /home/app
WORKDIR /home/app
COPY --from=builder /home/sfki /home/app
ENTRYPOINT /home/app/sfki
LABEL Name=sfki Version=0.0.2
EXPOSE 3000
