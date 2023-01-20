FROM golang:1.16-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN GOOS=linux go build -o restaurant-app .
CMD ["/app/restaurant-app"]