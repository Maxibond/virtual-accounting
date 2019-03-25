FROM golang

WORKDIR /go/src/app

COPY . .

RUN go build -v .

EXPOSE 8000

CMD ["./app"]