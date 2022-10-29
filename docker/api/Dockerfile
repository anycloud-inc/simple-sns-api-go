FROM golang:1.19-alpine

WORKDIR /api
COPY ./api .

RUN apk upgrade --update && apk --no-cache add git

# gccなど、ビルドやテストに必要なパッケージをインストール
RUN apk add --no-cache build-base

# ホットリロード用のライブラリをインストール
# RUN go install github.com/cosmtrek/air@v1.29.0
# RUN go get -u github.com/cosmtrek/air && go build -o /go/bin/air github.com/cosmtrek/air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# RUN go install github.com/cosmtrek/air@latest

# CMD ["air", "-c", ".air.toml"]