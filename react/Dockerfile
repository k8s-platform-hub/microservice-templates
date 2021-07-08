FROM node:8

# Install build essential and python
RUN apt-get update && apt-get install -y build-essential python

#Install deps
RUN mkdir /app
COPY app/package.json /app/package.json
RUN cd /app && npm install

#Add all source code
ADD app /app/
RUN cd /app && npm run build
RUN npm -g install serve

WORKDIR /app

#Default command
CMD ["serve", "-s", "build", "-p", "8080"]
