FROM golang:1.17
RUN apt-get update && apt-get install -y wget

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

WORKDIR /app

COPY go.mod .
COPY go.sum .

# dockerイメージの中にGo moduleを生成する
RUN go mod download

COPY . .

# -oオプションをつけると指定した位置に指定したファイル名で生成される
RUN go build

EXPOSE 8080
