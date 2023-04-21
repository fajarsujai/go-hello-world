# FROM golang:onbuild

# ARG BRANCH
# ARG PORT

# COPY .env.${BRANCH} /go/src/app/.env

# EXPOSE ${PORT}


# Create build stage based on buster image
FROM golang:1.16-buster AS builder


#build-arg
ARG BRANCH
ARG PORT

# Create working directory under /app
WORKDIR /app

# Copy over all go config (go.mod, go.sum etc.)
COPY go.* ./

# Install any required modules
RUN go mod download

# Copy over Go source code
COPY *.go ./

# Copy ENV
COPY .env.${BRANCH} .env

# Run the Go build and output binary under hello_go_http
RUN go build -o /hello_go_http

# Make sure to expose the port the HTTP server is using
EXPOSE ${PORT}

# Run the app binary when we run the container
ENTRYPOINT ["/hello_go_http"]