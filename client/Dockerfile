FROM golang:1.17.13
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main . 
CMD ["/app/main"]