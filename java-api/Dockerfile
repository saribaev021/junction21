


#FROM golang:1.17 as connec
#
#COPY ./app-api/connect /connect
#
#WORKDIR /connect
#
#RUN go build -o /wait /connect/main.go

FROM maven:3.8.3-jdk-11

#COPY --from=connec /wait /wait

COPY ./java-api/src /src

COPY ./java-api/pom.xml /pom.xml

COPY ./java-api/run.sh /run.sh

RUN mvn -f /pom.xml clean package

CMD ["sh", "/run.sh"]
