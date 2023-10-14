FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o godo ./main

EXPOSE 8080

CMD ["./godo"]