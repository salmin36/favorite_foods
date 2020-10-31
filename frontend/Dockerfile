FROM node:10.15.3-alpine

RUN mkdir /frontend
WORKDIR /frontend

COPY package.json ./
RUN npm install
COPY . ./
RUN npm run build
EXPOSE 8080
ENTRYPOINT npm run start
