FROM golang:1.8.5-jessie

# uncomment to install required debian packages
# add any package that is required after `build-essential`, end the line with \
# RUN apt-get update && apt-get install -y \
#     build-essential \
# && rm -rf /var/lib/apt/lists/*

# setup the working directory
WORKDIR /go/src/app

# add source code
ADD src src

# build the source
RUN go build src/main.go

# command to be executed on running the container
CMD ["./main"]

