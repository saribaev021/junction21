FROM golang:1.17
WORKDIR /go/src/KanatApp
COPY ./run.sh ./
RUN go get -u github.com/lib/pq
CMD ["sh", "run.sh"]
