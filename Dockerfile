FROM debian:bullseye

RUN mkdir -p /root/.kube/ 

RUN echo "fd97de6b91a121428112c52e5fe04a15" > /etc/machine-id

RUN apt-get update && \
    apt-get install -y apt-transport-https ca-certificates curl gnupg

RUN KUBECTL_VERSION=$(curl -L -s https://dl.k8s.io/release/stable.txt) && \
    curl -s "https://dl.k8s.io/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl" -o /usr/local/bin/kubectl && \
    chmod a+x /usr/local/bin/kubectl

RUN curl -sSfL https://raw.githubusercontent.com/kba-tools/besu-k8s-hooks/main/install.sh | sh -s
