#----- single-stage build----------------
# FROM golang:1.17.7-alpine
# RUN apk update && apk add git
# WORKDIR /go/src

# CMD ["go", "run", "main.go"]

#----- multi-stage build----------------
# 参考 https://zenn.dev/farstep/articles/a7819ec705dcc0
# 1.golang:1.17.7-alpineを使ってアプリケーションをbuildする
# 2.そこで生成されたバイナリファイルを軽量なイメージにコピーする

### Build stage ###
FROM golang:1.17.7-alpine AS builder
WORKDIR /app

# COPY src/ .
# srcをいきなりコピーするよりも、go.mod go.sumを先にコピーしておくほうがビルド時間が減る
# つまり、可能な限りビルドに必要なものだけを先にコピーし、依存関係をダウンロードする
COPY go.mod go.sum ./
RUN go mod download
COPY src/ .

# -o はアウトプットの名前を指定。(今回はmain)
# ビルドするファイル名を指定（今回は main.go）
RUN go build -o main /app/main.go
### Run stage ###
FROM alpine:latest
WORKDIR /app
# Build stageでビルドされたmainだけをRun stageにコピーする
COPY --from=builder /app/main .
CMD ["/app/main"]
