FROM golang:latest
WORKDIR /app

# add some necessary packages
ADD . /app


# prevent the re-installation of vendors at every change in the source code
COPY ./go.mod go.sum ./
RUN go mod download && go mod verify
RUN go build -o main .

# Install Compile Daemon for go. We'll use it to watch changes in go files
CMD ["./main"]
