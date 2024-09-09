FROM golang:alpine

WORKDIR /app

RUN apk update
RUN apk upgrade
RUN apk add --no-cache ffmpeg

COPY ./.. .

RUN go mod tidy
RUN go build -o application ./cmd/main.go
RUN chmod +x application

RUN mkdir tmp

EXPOSE 8182

CMD [ "./application" ]