package lib

import (
	"fmt"
	"os"
)

func SaveConfigTOML(folder string, networkId int, p2pHost string, p2pPort int) error {
	content := fmt.Sprintf(`# Network
network-id="%d"

# Gas -- Uncomment line below for gas free network
# min-gas-price=0

# Data
data-path="/data"
logging="INFO"

genesis-file="/config/genesis.json"
host-whitelist=["*"]

# P2P
p2p-host="%s"
p2p-port="%d"
max-peers=25

# RPC
rpc-http-enabled=true
rpc-http-host="0.0.0.0"
rpc-http-port=8545
rpc-http-api=["ADMIN", "DEBUG", "NET", "ETH", "MINER", "WEB3", "QBFT", "CLIQUE", "EEA", "IBFT"]
rpc-http-cors-origins=["all"]

# WS
rpc-ws-enabled=true
rpc-ws-host="0.0.0.0"
rpc-ws-port=8546
rpc-ws-api=["ADMIN", "DEBUG", "NET", "ETH", "MINER", "WEB3", "QBFT", "CLIQUE", "EEA", "IBFT"]

# GraphQL
graphql-http-enabled=true
graphql-http-host="0.0.0.0"
graphql-http-port=8547
graphql-http-cors-origins=["all"]

# Metrics
metrics-enabled=true
metrics-host="0.0.0.0"
metrics-port=9545

# Static Nodes
static-nodes-file="/config/static-nodes.json"
`,
		networkId,
		p2pHost,
		p2pPort)

	return os.WriteFile(fmt.Sprintf("%s/besu/config.toml", folder), []byte(content), 0644)
}
