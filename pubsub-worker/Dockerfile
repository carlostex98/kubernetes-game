FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
ENV GOOGLE_APPLICATION_CREDENTIALS=credentials2.json

COPY . .

RUN go build -o /pubsub-worker


CMD [ "/pubsub-worker" ]