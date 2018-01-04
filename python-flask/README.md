# python-flask

Boilerplate for a [Flask](http://flask.pocoo.org/) microservice.

## Add a python dependency

In order use new python package in your app, you can just add it to `src/requirements.txt` and the git-push or docker build process will
automatically install the package for you. If the `pip install` steps thorw some errors in demand of a system dependency,
you can install those by adding it to the `Dockerfile` at the correct place.

```
# src/requirements.txt:

flask
requests
gunicorn

# add your new packages one per each line
```

## Add a system dependency

The base image used in this boilerplate is [python:3](https://hub.docker.com/_/python/) debian. Hence, all debian packages are available for installation.
You can add a package by mentioning it in the `Dockerfile` among the existing `apt-get install` packages.

```dockerfile
# Dockerfile

FROM python:3

# install required debian packages
# add any package that is required after `python-dev`, end the line with \
RUN apt-get update && apt-get install -y \
    build-essential \
    python-dev \
&& rm -rf /var/lib/apt/lists/*

# install requirements
COPY src/requirements.txt /tmp/requirements.txt
RUN pip3 install -r /tmp/requirements.txt

# set /app as working directory
WORKDIR /app

# copy current directory to /app
COPY . /app

# run gunicorn server
# port is configured through the gunicorn config file
CMD ["gunicorn", "--config", "./conf/gunicorn_config.py", "src:app"]

```

## Deploy your existing Flask app

If you already have a Flask app and want to deploy it onto Hasura, read ahead:

- Replace the contents of `src/` directory with your own app's python files.
- Leave `k8s.yaml`, `Dockerfile` and `conf/` as it is.
- Make sure there is already a `requirements.txt` file present inside the new `src/` indicating all your python dependencies (see [above](#add-a-python-dependency)).
- If there are any system dependencies, add and configure them in `Dockerfile` (see [above](#add-a-system-dependency)).
- If the Flask app is not called `app`, change the last line in `Dockerfile` reflect the same.
  For example, if the app is called `backend`, the `CMD` line in `Dockerfile` will become:
  ```dockerfile
  CMD ["gunicorn", "--config", "./conf/gunicorn_config.py", "src:backend"]
  ```

## Debug

If the push fails with an error `Updating deployment failed`, or the URL is showing `502 Bad Gateway`/`504 Gateway Timeout`,
follow the instruction on the page and checkout the logs to see what is going wrong with the microservice:

```bash
# see status of microservice app
$ hasura microservice list

# get logs for app
$ hasura microservice logs app
```

## Local development

With Hasura's easy and fast git-push-to-deploy feature, you hardly need to run your code locally.
However, you can follow the steps below in case you have to run the code in your local machine.

### Without Docker

It is recommended to use a [Virtual Environment](http://docs.python-guide.org/en/latest/dev/virtualenvs/) for Python when you are running locally.
Don't forget to add these directories to `.gitignore` to avoid committing packages to source code repo.

```bash
# setup pipenv or virtualenv and activate it (see link above)

# go to app directory
$ cd microservices/app

# install dependencies
$ pip install -r src/requirements.txt

# Optional: set an environment variable to run Hasura examples 
# otherwise, remove Hasura examples, 
#   delete lines 5-8 from `src/__init__.py`
#   remove file `src/hasura.py`
$ export CLUSTER_NAME=[your-hasura-cluster-name]

# run the development server (change bind address if it's already used)
$ gunicorn --reload --bind "0.0.0.0:8080" src:app
```

Go to [http://localhost:8080](http://localhost:8080) using your browser to see the development version on the app.
You can keep the gunicorn server running and when you edit source code and save the files, the server will be reload the new code automatically.
Once you have made required changes, you can [deploy them to Hasura cluster](#deploy).

### With Docker

Install [Docker CE](https://docs.docker.com/engine/installation/) and cd to app directory:

```bash
# go to app directory
$ cd microservices/app

# build the docker image
$ docker build -t hello-python-flask-app .

# run the image with port bindings and CLUSTER_NAME environment variable
# as mentioned above, remove hasura.py if you don't want to add CLUSTER_NAME
$ docker run --rm -it -p 8080:8080 -e CLUSTER_NAME=[your-hasura-cluster-name] hello-python-flask-app

# app will be available at `http://localhost:8080`
# press Ctrl+C to stop the running container
```

For any change you make to the source code, you will have to stop the container, build the image again and run a new container.
If you mount the current directory as a volume, you can live-reload your code changes:

```bash
# go to app directory
$ cd microservices/app

# build the docker image
$ docker build -t hello-python-flask-app .

# run the container
$ docker run --rm -it -p 8080:8080 \
             -e CLUSTER_NAME=[your-hasura-cluster-name] \
             -v $(pwd):/app \
             hello-python-flask-app \ 
             gunicorn --reload --bind "0.0.0.0:8080" src:app
             
# app will be available at `http://localhost:8080`
# press Ctrl+C to stop the running container
```

Now, any change you make to your source code will be immediately updated on the running app.
