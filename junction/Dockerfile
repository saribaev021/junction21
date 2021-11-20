
FROM maven:3.8.3-jdk-11

COPY src /src

COPY pom.xml /pom.xml

RUN mvn -f /pom.xml clean package

CMD ["java", "-jar", "target/UserInfo-0.0.1-SNAPSHOT.jar" ]
