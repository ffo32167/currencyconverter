FROM golang:latest

WORKDIR /app

COPY ./ /app
#ENV PG_CONN_STR="host=postgres user=postgres password=pgpwd host=127.0.0.1:5432 dbname=postgres sslmode=disable"
ENV PG_CONN_STR="postgres://postgres:pgpwd@postgres:5432/postgres?sslmode=disable"
ENV PORT=":8080"
ENV CURRENCYFREAKS_CONN_STR="https://api.currencyfreaks.com/latest?apikey=f76751a9d86144a180db32b5de1a69a2&format=json"
#  CBR DATE FORMAT dd/mm/yyyy
ENV CBR_CONN_STR="https://www.cbr.ru/scripts/XML_daily.asp?date_req=02/03/2002"
ENV CURRENCIES="USD,EUR,RUB,JPY"
ENV CTX_TIMEOUT="1000"
ENV SOURCE="cbr"
ENV STORAGE="redis"
ENV REDIS_CONN_STR = "127.0.0.1:6379, ,0"

RUN go mod download

RUN go build -o currencyconverter cmd/http/*.go

CMD /app/currencyconverter