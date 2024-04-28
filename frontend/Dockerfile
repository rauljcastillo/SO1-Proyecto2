FROM node:20.11.0-alpine AS dist
WORKDIR /front
COPY package.json package-lock.json ./
RUN npm i

COPY . .
RUN npm run build

FROM nginx AS prod-stage

COPY --from=dist /front/dist /usr/share/nginx/html
EXPOSE 80:80
CMD ["nginx","-g","daemon off;"]