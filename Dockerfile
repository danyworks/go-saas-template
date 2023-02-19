FROM golang:1.19

ENV NAME $NAME

RUN apt-get update -yq && apt-get upgrade -yq

USER root

WORKDIR /data

COPY . .

#RUN go mod download

RUN go build -o ./bin/marketplace ./cmd/backend/main.go 

EXPOSE 8080

CMD [ "/data/bin/marketplace","8080" ]