FROM golang:1.23.0

RUN apt update

ENV GO111MODULE on
WORKDIR /go/src/work

COPY ./../go.mod ./go.mod
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/xo/xo@latest
RUN go install github.com/amacneil/dbmate@latest
RUN go install golang.org/x/tools/cmd/goimports@latest
RUN go mod tidy
RUN export DATABASE_URL=mysql://root:root@db:3306/test