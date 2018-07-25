FROM mhart/alpine-node:9
WORKDIR /src
# Add app source files
ADD app/src /src
#install node modules
RUN npm install --only=production

CMD ["node", "server.js"]
