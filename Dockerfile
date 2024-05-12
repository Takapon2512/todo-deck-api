# GoのAlpineベースイメージを使用
FROM golang:1.21-alpine

# アプリケーションのワーキングディレクトリを設定
WORKDIR /app

# パッケージリストの更新とシステムパッケージのアップグレード
RUN apk update && apk upgrade

# Gitをインストール
RUN apk add git

# go.modとgo.sumをホストからコンテナにコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをホストからコンテナにコピー
COPY . .

# CMD命令を使用して、アプリケーションを起動
CMD ["go", "run", "main.go"]
