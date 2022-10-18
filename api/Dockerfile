FROM golang:1.19-alpine

WORKDIR /go/src
COPY ./api/src .

RUN apk upgrade --update && apk --no-cache add git

# gccなど、ビルドやテストに必要なパッケージをインストール
RUN apk add --no-cache build-base

# ホットリロード用のライブラリをインストール
# RUN go get github.com/cosmtrek/air@v1.29.0
RUN go get -u github.com/cosmtrek/air && go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]