#docker image rm -f chatapp
#docker rm -f chatApp_backend
#docker build -t chatapp -f Dockerfile .
#docker run -p 9090:9090 --name chatApp_backend --network chatnet -v /home/rosy/chatAppFile:/app/chatAppFile chatapp
FROM golang:alpine AS build

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN go env -w GOPROXY=https://goproxy.io
RUN go mod download

COPY *.go ./
RUN go build -o /chatapp

FROM alpine:latest

WORKDIR /app

COPY --from=build /chatapp /app/chatapp

EXPOSE 9090

#这个必须要加，不然没有权限

ENTRYPOINT ["/app/chatapp"]

