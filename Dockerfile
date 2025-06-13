# Stage 1: Build
FROM golang:1.24.4-alpine AS build

WORKDIR /ocrolus-app

RUN apk add --no-cache make git

# Install goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN make build

# Stage 2: Run
FROM alpine:latest

WORKDIR /bin/app

RUN apk add --no-cache postgresql-client

COPY --from=build /ocrolus-app/.env /bin/app/.env
COPY --from=build /go/bin/goose /usr/local/bin/goose
COPY --from=build /ocrolus-app/sql/migrations /bin/app/sql/migrations
COPY --from=build /ocrolus-app/scripts/entrypoint.sh /bin/app/entrypoint.sh
RUN chmod +x /bin/app/entrypoint.sh
COPY --from=build /ocrolus-app/bin/app/ocrolus-task /bin/app/ocrolus-task

EXPOSE 8080
