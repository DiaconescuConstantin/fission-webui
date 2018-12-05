FROM node:alpine as nodebuilder

RUN mkdir -p /app
WORKDIR '/app'
COPY ./package.json ./
RUN yarn install
COPY . .
RUN yarn run build
RUN ls -lsa -R dist/

FROM golang:1.11.2 as gobuilder

RUN mkdir -p /go/src/go-server
WORKDIR /go/src/go-server/
COPY ./go-server/server.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=nodebuilder /app/dist ./static
COPY --from=gobuilder /go/src/go-server/server .
EXPOSE 3000
CMD ["./server"]