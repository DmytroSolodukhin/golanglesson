FROM golang:latest
WORKDIR /app
COPY .. ./
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -a -o ./bin/svc
EXPOSE 8080
RUN go build /main.go
ENTRYPOINT ["main"]
