FROM golang:1.16
WORKDIR /app
ADD . .
RUN go build -o main .
CMD /app/main
