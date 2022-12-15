##Stage 1
## Start from the latest golang base image for builder
FROM golang:alpine as builder
# ENV GO111MODULE=on
RUN mkdir /go/src/service
ADD . /go/src/service
WORKDIR /go/src/service
RUN apk add git
RUN go env -w GOINSECURE=192.168.223.213
RUN go mod tidy

# BUILD DOCS SWAGGER
# RUN go get -u github.com/swaggo/swag/cmd/swag 
RUN go install -v github.com/swaggo/swag/cmd/swag@latest
# RUN go install -v github.com/swaggo/swag/cmd/swag@v1.7.4
# &&  go get -u github.com/swaggo/http-swagger && go get -u github.com/alecthomas/template
RUN swag init -g main.go -o Library/Swagger/docs
ENV GOOGLE_APPLICATION_CREDENTIALS "/go/src/service/Infrastructures/GCP/secret-manager-key.json"
#TESTING
#RUN CGO_ENABLED=0 GOOS=linux go test ./Controller/...

# BUILD ENGINE
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -s" -o main main.go
RUN chmod 755 /go/src/service/main

##Stage 2
## Start from the latest alpine base image
FROM alpine:latest
LABEL maintainer="Muhammad Nasrul <mnasruul@gmail.com>"
RUN addgroup -S nasrul && adduser -S nasrul -G nasrul
RUN apk update && apk upgrade && apk add --no-cache tzdata
ARG APP_ENV
ENV env_state=$APP_ENV

RUN mkdir -p /app
ENV GOOGLE_APPLICATION_CREDENTIALS "/go/src/service/Infrastructures/GCP/secret-manager-key.json"
WORKDIR /app
# COPY config.yml ./
COPY --chown=nasrul:nasrul --from=builder /go/src/service/Environment /go/src/service/Environment
COPY --chown=nasrul:nasrul --from=builder /go/src/service/Infrastructures/GCP /go/src/service/Infrastructures/GCP
COPY --chown=nasrul:nasrul --from=builder  /go/src/service/main /app 
USER nasrul

# Expose port 7000 to the outside world.
EXPOSE 7000

CMD /app/main -env=$env_state

