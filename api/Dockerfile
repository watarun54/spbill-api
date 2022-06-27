FROM golang:1.18.3-alpine

# ログに出力する時間をJSTにするため、タイムゾーンを設定
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

ENV ROOT=/go/src/github.com/watarun54/spbill-api/server
WORKDIR ${ROOT}

# ModuleモードをON
ENV GO111MODULE=on

COPY . .

RUN go mod tidy && \
    go install github.com/cosmtrek/air@latest

CMD ["air"]
