FROM golang:alpine

WORKDIR /app

COPY . .

RUN apk update && apk add gcc g++ sqlite

CMD ["go", "run", "main.go"]