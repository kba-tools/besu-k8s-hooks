# ==============================
# 🧩 Stage 1: Builder
# ==============================
FROM debian:bullseye AS builder

# Install build dependencies
RUN apt-get update && \
    apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    jq && \
    rm -rf /var/lib/apt/lists/*

# Set machine ID (optional)
RUN echo "fd97de6b91a121428112c52e5fe04a15" > /etc/machine-id

# Install kubectl (latest stable)
RUN KUBECTL_VERSION=$(curl -L -s https://dl.k8s.io/release/stable.txt) && \
    curl -LO "https://dl.k8s.io/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl" && \
    install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl && \
    rm kubectl

# Create a safe working directory
WORKDIR /opt/besu

# Install Besu hooks
RUN curl -sSfL https://raw.githubusercontent.com/kba-tools/besu-k8s-hooks/main/install.sh | sh -s -- -b ./hooks

# ==============================
# 🧩 Stage 2: Runner
# ==============================
FROM debian:bullseye-slim AS runner

# Install minimal runtime dependencies
RUN apt-get update && \
    apt-get install -y \
    ca-certificates \
    curl \
    jq && \
    rm -rf /var/lib/apt/lists/*

# Copy necessary files from builder
COPY --from=builder /usr/local/bin/kubectl /usr/local/bin/kubectl
COPY --from=builder /opt/besu/hooks /opt/besu/hooks

# Optional: create kube config directory and machine ID
RUN mkdir -p /root/.kube && \
    echo "fd97de6b91a121428112c52e5fe04a15" > /etc/machine-id

WORKDIR /opt/besu

# Default shell
CMD ["/bin/bash"]
