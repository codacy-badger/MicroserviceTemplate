################# Start build stage #################
# Base on golang image
FROM golang:1.14 AS builder

WORKDIR /app

# Copy the mod files into image
COPY . .

# Build the module and run tests
RUN make build
RUN make test

################# Start deploy stage ################
# Base on alpine image
FROM alpine AS deploy

# Copy the build artifact
COPY --from=builder /app/main .
COPY --from=builder /app/docs .

# Expose port used by the app
EXPOSE 8080

# Make image an executable
ENTRYPOINT [ "./main" ]
