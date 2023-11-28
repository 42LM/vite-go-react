# Build the Go API
FROM golang:alpine AS builder
ADD ./server /app/server
WORKDIR /app/server
RUN go mod download
RUN go install github.com/mattn/go-sqlite3
RUN apk add --update gcc musl-dev
RUN CGO_ENABLED=1 go build -ldflags "-w" -a -o /main .

# Build the React application
FROM node:alpine AS node_builder
ADD ./client /client
WORKDIR /client
RUN npm install
RUN npm run build

# Final stage build, this will be the container
# that we will deploy to production
FROM alpine:latest
ADD ./db /db
RUN apk --update-cache add sqlite \
    && rm -rf /var/cache/apk/* \
    && ./db/createdb.sh
RUN apk add --no-cache libc6-compat
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
COPY --from=node_builder /client/dist ./web
RUN chmod +x ./main
EXPOSE 8080
CMD ["./main"]
