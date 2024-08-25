# dev
FROM golang:1.22-bullseye AS dev
WORKDIR /work/CanBook-backend-go

# Install Air for live reload
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /go/bin

# Copy go.mod and go.sum to leverage Docker cache for dependency download
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY ./ ./

# Build the binary
RUN make build-linux

# release
FROM alpine AS release
RUN apk --no-cache add tzdata

# Set the timezone
ENV TZ=Asia/Tokyo

# Copy built binary from the dev stage
COPY --from=dev /work/CanBook-backend-go/build/CanBook-backend-go-linux-amd64 /usr/local/bin/CanBook-backend-go

# Ensure the binary is executable
RUN chmod +x /usr/local/bin/CanBook-backend-go

# Expose the application port
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/usr/local/bin/CanBook-backend-go"]
