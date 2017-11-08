# quickstart-docker-git

[Hasura](https://hasura.io) is a platform for building and deploying application backends. The platform is a [Kubernetes](https://kubernetes.io) based PaaS (*Platform-as-a-Service to deploy backends*) and a PostgreSQL based BaaS (*Backend-as-a-Service to rapidly build backends*).

To help you quickly bootstrap your microservice in a Hasura project, this repository maintains a collection of starter kits for popular options.

Some examples:
- Hello-world for [python-flask](https://github.com/hasura/quickstart-docker-git/tree/master/python-flask), [java-spring boot](https://github.com/hasura/quickstart-docker-git/tree/master/java-spring-boot), [nodejs-express](https://github.com/hasura/quickstart-docker-git/tree/master/nodejs-express) [php-apache](https://github.com/hasura/quickstart-docker-git/tree/master/php-apache), [.Net](https://github.com/hasura/quickstart-docker-git/tree/master/csharp-aspnet), etc.
- Static file service based on [nginx](https://github.com/hasura/quickstart-docker-git/tree/master/nginx)
- Web service in Swift ([Vapor](https://github.com/hasura/quickstart-docker-git/tree/master/swift-vapor), [Perfect](https://github.com/hasura/quickstart-docker-git/tree/master/swift-perfect))
- Deploying a [MySQL](https://github.com/hasura/quickstart-docker-git/tree/master/mysql) service
- [Wordpress](https://github.com/hasura/quickstart-docker-git/tree/master/wordpress)

**What is not included**: This is just a collection of starter kits. The idea is to modify the kits as you need to get to a basic setup that works.

# Getting Started in 5 easy steps

Make sure you've installed the [`hasura CLI`](https://docs.hasura.io/0.15/manual/install-hasura-cli.html) tool.

To use any of these templates, follow these instructions:

```bash
# Step 1: Create a new hasura project if you don't have one already
$ hasura quickstart hasura/base my-project

# Step 2: cd inside the project directory and generate
# the microservice folder (called app) based on a template
$ cd my-project
$ hasura microservice generate app --template=nodejs-express

# Step 3: Configure your project so that you can deploy
# the microservice with a 'git push hasura master'.
# The command below will add the configuration to the
# right file.
$ hasura conf generate-remote app >> conf/ci.yaml

# Step 4: Configure a route (subdomain or path) on which
# you want to expose this microservice externally/publicly
$ hasura conf generate-route app >> conf/routes.yaml

# Step 5: Commit and push!
$ git add conf/ci.yaml conf/routes.yaml microservices/app
$ git commit -m 'Adds the app microservice'
$ git push hasura master

# FINISH: You're all done! Run the command below to
# list all your microservices and their URLs.
$ hasura microservices list
```

-------------

# Contribution Needed!

Please fork, file comments/bugs and submit your PRs!  We've created a list of
issues where active help is required:
[help-wanted](https://github.com/hasura/quickstart-docker-git/issues?q=is%3Aissue+is%3Aopen+label%3Ahelp-wanted).

Specifically,

1. Add support for more frameworks
2. Optimise ``Dockerfile``s for faster builds, or for production (esp. for compiled languages)

Some important things to keep in mind when contributing:

1. Expose only one port, and one data volume to help keep things simple
2. Annotate the ``Dockerfile`` with comments where you expect users to modify
3. Try to document the following major use cases when writing your README: ``docker build``, ``git push`` based automated docker build, ``docker run``
