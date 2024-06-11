FROM node:20.10.0

WORKDIR /app

RUN npm install -g @typespec/compiler

COPY package*.json ./

RUN npm install
