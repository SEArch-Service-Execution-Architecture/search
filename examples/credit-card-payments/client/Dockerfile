# Use an official Maven runtime as a parent image
FROM maven:3-eclipse-temurin-24-alpine AS build

# Set the working directory to /app
WORKDIR /app

# WARNING! Paths in the Dockerfile are relative to the root of the repository.
# We do that to be able to copy the generated stubs from the gen directory.
# The context is set in the docker-compose.yml file.
# COPY examples/credit-card-payments/client/pom.xml .
# RUN mvn verify --fail-never
COPY examples/credit-card-payments/client/ .
RUN rm lib contract.fsa
COPY examples/credit-card-payments/contract.fsa .
COPY gen/java ./lib

# Compile the project with Maven
RUN mvn package
CMD ["java", "--enable-native-access=ALL-UNNAMED", "-Dio.grpc.NameResolverProvider=dns", "-jar", "/app/target/clientapp-1.0-SNAPSHOT-jar-with-dependencies.jar"]

FROM eclipse-temurin:24-alpine

# Set the working directory to /app
WORKDIR /app

COPY examples/credit-card-payments/contract.fsa .
# Copy the compiled JAR file from the build container to the runtime container
COPY --from=build /app/target/clientapp-1.0-SNAPSHOT-jar-with-dependencies.jar ./clientapp.jar

# Run the JAR file with necessary Java flags for Java 21+
CMD ["java", "--enable-native-access=ALL-UNNAMED", "-Dio.grpc.NameResolverProvider=dns", "-jar", "clientapp.jar"]
