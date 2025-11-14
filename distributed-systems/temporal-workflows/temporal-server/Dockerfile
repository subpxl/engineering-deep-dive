# Use a minimal base image
FROM ubuntu:22.04

# Install required dependencies
RUN apt-get update && apt-get install -y \
    ca-certificates \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Copy the local temporal binary into the image
COPY temporal /usr/local/bin/temporal

# Make it executable
RUN chmod +x /usr/local/bin/temporal

# Expose Temporal default ports
EXPOSE 7233 8233 9100

# Run Temporal in dev mode with metrics
CMD ["/usr/local/bin/temporal", "server", "start-dev", "--port", "7233", "--ui-port", "8233", "--ui-ip", "0.0.0.0", "--ip", "0.0.0.0", "--metrics-port", "9100"]
