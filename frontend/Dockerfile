FROM node:10

RUN mkdir /frontend
WORKDIR /frontend


COPY package.json ./
COPY dist ./dist
COPY src ./src

RUN ["find", "/frontend"]
EXPOSE 8080
ENTRYPOINT  ["node", "/frontend/src/server.js"]   
#ENTRYPOINT ["/bin/bash"] 