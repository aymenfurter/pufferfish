FROM golang:1.17

WORKDIR /app

COPY memory.go .

RUN go build -o pufferfish memory.go

CMD [ "./pufferfish" ]
