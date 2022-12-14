FROM golang:1.19-alpine

# Create a directory for the app
RUN mkdir /app
 
# Copy all files from the current directory to the app directory
COPY . /app
 
# Set working directory
WORKDIR /app

# Set prod config env variables
ENV APP_MODE=prod \
    PORT=3001

# go build will build an executable file named server in the current directory
RUN go build -o ./cronApi ./cmd/api

# Run the server executable
CMD [ "/app/cronApi" ]