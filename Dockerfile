FROM golang:alpine AS go
RUN apk add build-base
WORKDIR /backend

COPY /backend/swagger.sh .
RUN ./swagger.sh install

COPY /backend/go.mod .
RUN go mod download

COPY ./backend .
RUN ./swagger.sh init
RUN go build -o app

FROM alpine AS logo
RUN apk add figlet
WORKDIR /logo

RUN figlet Café > logo.txt

FROM node:latest AS vue
WORKDIR /frontend
COPY ./frontend/package.json .
COPY ./frontend/package-lock.json .
RUN npm install

COPY --from=go /backend/docs/ /backend/docs/
COPY ./frontend .
RUN npm run types:openapi -i /backend/docs/swagger.json -o src/services/openapi
RUN npm run build --verbose

FROM alpine AS final
WORKDIR /app
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh
ENV GIN_MODE=release

COPY --from=logo /logo/logo.txt .
COPY --from=go /backend/app .
COPY --from=go /backend/config.toml .
COPY --from=vue /frontend/dist/ ./templates/

ENTRYPOINT ["/app/entrypoint.sh"]
