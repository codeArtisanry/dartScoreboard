FROM node:16.13-alpine

WORKDIR /app

COPY . ./

RUN npm install

EXPOSE 3000

RUN npm run build

RUN cp .env.docker .env

CMD [ "npm", "start" ]
