FROM golang:1.19-alpine AS build
RUN apk update
WORKDIR /app
COPY . /app
RUN go build

FROM alpine:3.17 AS prod
ENV PORT=8000 \
    REDIS_HOST=localhost \
    REDIS_PORT=6379 \
    REDIS_PASSWORD=
EXPOSE 3000
COPY --from=build /app/backend /bin/backend
CMD ["/bin/backend"]
