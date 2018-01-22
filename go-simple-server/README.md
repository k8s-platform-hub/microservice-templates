# go-simple-server 

A simple golang web server with basic logging, tracing, health check, graceful shutdown and zero dependencies. This microservice is based on a [GitHub gist](https://gist.github.com/enricofoltran/10b4a980cd07cb02836f70a4ab3e72d7) by [enricofoltran](https://github.com/enricofoltran), featured on [HackerNews](https://news.ycombinator.com/item?id=16090977) which was later moved to [enricofoltran/simple-go-server](https://github.com/enricofoltran/simple-go-server).

## Local development

### With Docker

- Install [Docker CE](https://docs.docker.com/engine/installation/)

```bash
# go to app directory
$ cd microservices/app

# build the docker image
$ docker build -t simple-go-server-app .

# run the image
$ docker run --rm -it -p 5000:5000 simple-go-server-app

# app will be available at http://localhost:5000
# press Ctrl+C to stop the server
```

### Without Docker

- Install [Golang](https://golang.org/doc/install)
- Move the `simple-go-server` directory to your `GOPATH` and cd into the directory

```bash
# change to app directory
$ cd mircoservices/app

# run the local server script
# windows users: run this in git bash
$ go run src/main.go

# app will be available at http://localhost:5000
# press Ctrl+C to stop the server
```
