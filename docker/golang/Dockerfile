# 開発環境用

FROM golang:1.20.3-alpine

RUN apk update
RUN apk add alpine-sdk

WORKDIR /app/go

# ホストマシンからソースコードをコピー
COPY ./go-app .

# go.modからmodをダウンロードする
RUN go mod download

# air のインストール
# 開発環境用でリアルタイムにコンパイルさせる
RUN go install github.com/cosmtrek/air@latest

#go install コマンドによって air がインストールされた場所（/root/go/bin）を PATH に追加しています。
# これにより、air コマンドがコンテナ内のどこからでも実行できるようになります。
ENV PATH=$PATH:/root/go/bin

# # アプリケーションをビルド（本番環境用）
# RUN go build -o /app/main .
# # 実行コマンド
# CMD [ "/app/main" ]
