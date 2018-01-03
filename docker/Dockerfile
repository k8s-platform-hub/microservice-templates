# Define the base image
FROM alpine:3.5

# Set environment variables
ENV APP_HOME /app

# Create app directory
RUN mkdir $APP_HOME

# Copy files from dev environment to container
COPY bin/app $APP_HOME

# Give exec permissions
RUN chmod +x $APP_HOME/app

# Set directory path
WORKDIR $APP_HOME

# Run executable
CMD ["./app"]
