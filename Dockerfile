FROM golang:rc-bullseye
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o server cmd/api/main.go
CMD [ "/app/server" ]
