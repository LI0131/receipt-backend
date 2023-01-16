FROM golang:bullseye as builder

RUN apt-get update && \
    apt-get -y install ca-certificates \
                       libtesseract-dev \
                       tesseract-ocr \
                       tesseract-ocr-eng

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go build -o receipt-api cmd/receipt-backend-api/main.go


FROM debian:bullseye-slim as app

WORKDIR /

RUN apt-get update && \
    apt-get -y install ca-certificates \
                       libtesseract-dev \
                       tesseract-ocr \
                       tesseract-ocr-eng

RUN mkdir /static

COPY --from=builder /go/src/app/receipt-api ./receipt-api

CMD ["./receipt-api"]
