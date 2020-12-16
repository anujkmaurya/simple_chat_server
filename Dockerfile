FROM golang:1.15.4

ENV CGO_ENABLED 0

WORKDIR $GOPATH/src/

COPY . .

# Download all the dependencies
RUN go get -d -v ./...

RUN go install -v ./...

# This container exposes port 9000 to the outside world
EXPOSE 9000

# Run the executable
CMD ["server"]

