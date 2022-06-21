FROM golang:1.14

WORKDIR /app

COPY . .

RUN go mod download

RUN go build

EXPOSE 3000

CMD ./zeetgo2