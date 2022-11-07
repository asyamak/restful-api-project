##source image of golang base image
FROM golang:1.19-alpine as build
#maintainer info
LABEL maintainer = "asyamak"

# Install git.
# Git is required for fetching the dependencies.
#RUN apk update && apk add --no-cache git

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

#RUN apt-get update && go get github.com/pressly/goose/cmd/goose
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base
ENV CGO_ENABLED=0
RUN apk update
RUN apk add --no-cache git gcc
RUN  go get github.com/pressly/goose/cmd/goose
# RUN go install github.com/pressly/goose/v3/cmd/goose@latest
# Download all the dependencies
RUN go get -d -v ./...

# RUN goose -dir migration postgres "postgresql://postgres:5432/goose?sslmode=disable" up
# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Install the package
RUN go install -v ./... 
# COPY 20221106144353_rest_api.sql ./
# RUN go get github.com/jmoiron/sqlx
# RUN go get -v github.com/lib/pq
# RUN go get -v github.com/DATA-DOG/go-sqlmock


# Build the Go app
RUN go build -o main ./main.go

# Expose port 8080 to the outside world
EXPOSE 4040

# Run the executable
ENTRYPOINT ["/app/main"]