FROM golang:1.19-alpine3.16 As builder

RUN apk add --no-cache bash

WORKDIR /app

COPY . .

RUN go build -ldflags="-s" -o url_shortener

FROM alpine
WORKDIR /app
COPY --from=builder /app/url_shortener /app

CMD [ "./url_shortener" ]