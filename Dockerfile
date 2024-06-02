FROM golang:1.22-alpine as builder

WORKDIR /app
COPY go.* .
RUN go mod download
COPY . .
RUN go build

FROM restic/restic:0.16.4
RUN apk add --no-cache rclone bash curl docker-cli
COPY --from=builder /app/autorestic /usr/bin/autorestic
ENTRYPOINT []
CMD [ "bash" ]
