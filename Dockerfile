
#build stage
FROM golang:alpine AS builder
RUN mkdir -p /go/src/github.com/loadfield/sfki
WORKDIR /go/src/github.com/loadfield/sfki
COPY . /go/src/github.com/loadfield/sfki
# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache git
RUN go get -v gopkg.in/yaml.v2
RUN go get -v github.com/go-chi/chi
RUN go get -v github.com/graphql-go/graphql
RUN go build -ldflags "-s -w"

#final stage
FROM alpine:latest
RUN mkdir -p /home/app
WORKDIR /home/app
COPY --from=builder /go/src/github.com/loadfield/sfki /home/app
ENTRYPOINT /home/app/sfki
LABEL Name=sfki Version=0.0.1
EXPOSE 3000
