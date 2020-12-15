# 使用するGolangのイメージを指定する
FROM golang:1.12.6

ENV GO111MODULE=on


# ワーキングディレクトリを指定する
WORKDIR /go/src

COPY docker/golang/go.mod .
COPY docker/golang/go.sum .


# 必要なパッケージなどをインストールする

RUN apt-get update -qq && \ 
    apt-get install -y git mysql-client vim


RUN go mod download


# golang直下のディレクトリをコンテナ上に載せる
ADD docker/golang/ .




RUN go build


