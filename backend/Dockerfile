FROM golang:1.19-alpine AS build
RUN apk update
WORKDIR /app
COPY . /app
RUN go build

FROM alpine:3.17 AS prod
ENV PORT=8080 \
    REDIS_HOST=pepe \
    REDIS_PORT=1234
COPY --from=build /app/backend /bin/backend
CMD ["/bin/backend"]
