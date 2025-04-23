#!/bin/bash

# Stop and remove any existing container with the same name
docker rm -f simple-grpc-test 2>/dev/null

# Run the container in the background, mapping port 50051
docker run --name simple-grpc-test -d -p 50051:50051 simple-grpc

echo "Container started. Waiting 2 seconds for server to initialize..."
sleep 2

echo "Testing 'ListConnections' API..."
# Run grpcurl from inside the container to test ListConnections
docker exec simple-grpc-test grpcurl -plaintext -proto cb.proto -import-path / -d '{}' localhost:50051 cb.Greet/ListConnections

# Keep the container running in case you want to do more tests
echo -e "\nContainer is still running. Use 'docker rm -f simple-grpc-test' to stop and remove it when done."