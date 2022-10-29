FROM golang:1.19-alpine

WORKDIR /api
COPY ./api .

RUN apk upgrade --update && apk --no-cache add git

# gccなど、ビルドやテストに必要なパッケージをインストール
RUN apk add --no-cache build-base

# ホットリロード用のライブラリをインストール
RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]