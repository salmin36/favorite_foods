FROM node:10

RUN mkdir /frontend
WORKDIR /frontend

COPY package.json ./
RUN npm install
COPY . ./
RUN npm run build
EXPOSE 8080
RUN npm run start
#ENTRYPOINT  ["node", "/frontend/src/server.js"]   
#ENTRYPOINT ["/bin/bash"] 