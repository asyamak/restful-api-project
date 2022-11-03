##source image of golang base image
FROM golang:1.19-alpine as builder
#maintainer info
LABEL maintainer = "asyamak"

# Install git.
# Git is required for fetching the dependencies.
#RUN apk update && apk add --no-cache git

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Download all the dependencies
RUN go get -d -v ./...

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Install the package
RUN go install -v ./... 

# RUN go get github.com/jmoiron/sqlx
# RUN go get -v github.com/lib/pq
# RUN go get -v github.com/DATA-DOG/go-sqlmock


# Build the Go app
RUN go build -o main ./main.go

# Expose port 8080 to the outside world
EXPOSE 4040

# Run the executable
ENTRYPOINT ["/app/main"]