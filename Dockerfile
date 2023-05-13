FROM golang:1.20-alpine AS build

ENV HTTPS_PROXY="http://fodev.org:8118"

WORKDIR /app

COPY . ./

# Install dependencies
RUN go mod download && \
  # Build the app
  GOOS=linux GOARCH=amd64 go build -o main && \
  # Make the final output executable
  chmod +x ./main

FROM alpine:latest

# Install os packages
RUN apk --no-cache add bash

WORKDIR /app

COPY --from=build /app/main .

CMD ["./main"]

EXPOSE 8000