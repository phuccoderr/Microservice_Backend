FROM openjdk:17-jdk-alpine

RUN mkdir /app

WORKDIR /app

COPY target/review-service-0.0.1-SNAPSHOT.jar review-service-0.0.1-SNAPSHOT.jar

ENTRYPOINT ["java", "-jar", "review-service-0.0.1-SNAPSHOT.jar"]