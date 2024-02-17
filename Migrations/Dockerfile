# Use the official migrate image as base image
FROM migrate/migrate

# Set the working directory to /migrations
WORKDIR /migrations

# Copy the migration files from the host to the container
COPY . /migrations

# Define the entrypoint and default command for the container
#ENTRYPOINT ["/bin/sh", "-c"]

#CMD ["migrate -database "postgres://postgres:helloworld@db:5432/postgres?sslmode=disable" -path /migrations force 1 && yes | migrate -path /migrations -database "postgres://postgres:helloworld@db:5432/postgres?sslmode=disable" down && migrate -path /migrations -database "postgres://postgres:helloworld@db:5432/postgres?sslmode=disable" up"]
