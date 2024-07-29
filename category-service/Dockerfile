FROM openjdk:17-jdk-alpine

RUN mkdir /app

WORKDIR /app

COPY target/category-service-0.0.1-SNAPSHOT.jar category-service-0.0.1-SNAPSHOT.jar

ENTRYPOINT ["java", "-jar", "category-service-0.0.1-SNAPSHOT.jar"]