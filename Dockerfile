FROM golang:1.18.2-bullseye as deploy-builder

# リリース用のビルドを行うコンテナ
WORKDIR /app
copy go.mod go.sum ./
RUN go mod donwload

Copy . . 
RUN go build -trimpath -ldflags  "-w -s" -o app

FROM debian:bullseye-slim as deploy

# マネージドサービス上で鵜がかす事を想定したリリース用のコンテナイメージを作成するステージ
RUN apt-get update

Copy --from=deploy-builder /app/app .
CMD ["./app"]

FROM golang:1.18.2 as dev
# ローカルで開発する時に利用するコンテナイメージを作成するステージ
WORKDIR /app
#ホットリロード開発を使用するためのライブラり
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]