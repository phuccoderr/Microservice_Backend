FROM openjdk:17-jdk-alpine

RUN mkdir /app

WORKDIR /app

COPY target/discovery-server-0.0.1-SNAPSHOT.jar discovery-server-0.0.1-SNAPSHOT.jar

ENTRYPOINT ["java", "-jar", "discovery-server-0.0.1-SNAPSHOT.jar"]
