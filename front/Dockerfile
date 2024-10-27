# Etapa 1: Construcción de la aplicación Node.js
FROM cgr.dev/chainguard/node as build

#ENV NODE_ENV=production

COPY --chown=node:node ["./", "/app/"]


WORKDIR /app

RUN npm install

RUN npm run build

RUN npm install -g serve

EXPOSE 8080

CMD ["serve -s build -l 8080"]
