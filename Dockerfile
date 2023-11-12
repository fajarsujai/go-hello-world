# Create build stage based on buster image
FROM golang:1.16-buster AS builder


#build-arg
ARG BRANCH
ARG PORT
ARG PROJ_NAME

# Create working directory under /app
WORKDIR /app

# Copy over all go config (go.mod, go.sum etc.)
COPY go.* ./

# Install any required modules
RUN go mod download

# Copy over Go source code
COPY *.go ./

# Copy ENV
COPY ./.env ./.env

# Run the Go build and output binary under hello_go_http
RUN go build -o /hello_go_http

# Make sure to expose the port the HTTP server is using
EXPOSE 8000

# Run the app binary when we run the container
ENTRYPOINT ["/hello_go_http"]