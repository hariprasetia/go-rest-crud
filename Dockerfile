FROM golang:1.16-alpine as builder

ENV GO111MODULE=on

LABEL maintainer="Hari Prasetia <hari.prstya@gmail.com>"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/config/.env ./config/.env

EXPOSE 3000

CMD ["./main"]