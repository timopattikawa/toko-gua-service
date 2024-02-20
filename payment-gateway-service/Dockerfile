FROM golang:1.16-alpine

LABEL author="Timo Pattikawa"

WORKDIR /app

RUN go build -o toko-gue-payment-service

EXPOSE 3000

CMD [ "./toko-gue-payment-service" ]