FROM mhart/alpine-node:7.10

# APK is the alpine-pakage-manager Python gcc is need for node-sass
RUN apk add --no-cache make gcc g++ python

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
