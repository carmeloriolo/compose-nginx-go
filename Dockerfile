##
## Build
##
FROM golang:1.16-alpine AS build
ARG AppName

WORKDIR /monorepo

COPY ./ ./
RUN go mod download

COPY ./services/${AppName}/*.go ./

RUN go build -o /app

##
## Deploy
##
FROM alpine:3.14
ARG AppName
EXPOSE 8080

WORKDIR /

COPY --from=build /app /app

ENTRYPOINT ["/app"]