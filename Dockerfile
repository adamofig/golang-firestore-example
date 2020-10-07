FROM golang:1.14.0-alpine3.11
COPY . /app
WORKDIR /app
RUN cd /app && go build -o goapp
EXPOSE 8080
ENTRYPOINT ./goapp
