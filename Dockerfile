FROM golang:1.22.5-alpine as build
WORKDIR /app
COPY go.mod go.sum .
RUN go mod tidy
RUN GOOS=linux go build -o ./bin/dump_project ./cmd/main.go

FROM alpine:latest
WORKDIR /usr/app
COPY --from-build /app/bin/dump_project ./bin/dump_project
EXPOSE 8001

CMD [ "./bin/dump_project" ]
