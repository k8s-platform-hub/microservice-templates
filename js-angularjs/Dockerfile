FROM node:6

#Install deps
RUN npm install -g http-server grunt-cli bower yo generator-karma generator-angular
RUN mkdir /app
COPY package.json /app/package.json
COPY bower.json /app/bower.json

WORKDIR /app

RUN npm install
RUN bower --allow-root install 

#Add all source code
ADD . . 

# Execute build
RUN grunt build 

# Make command executable
RUN chmod +x /app/runserver.sh

# Default command
CMD ["/app/runserver.sh"]
