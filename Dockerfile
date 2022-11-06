FROM node AS node
WORKDIR /app

COPY ui /app/ui
WORKDIR /app/ui
RUN npm install
RUN npm run build

FROM golang AS golang
WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/
COPY main.go /app/
RUN mkdir /app/ui
COPY --from=node /app/ui/build /app/ui/build

RUN go mod vendor
RUN go build -o main main.go

FROM alpine
RUN apk update
RUN apk add gcompat
WORKDIR /app
COPY --from=golang /app/main /app/

CMD [ "/app/main" ]