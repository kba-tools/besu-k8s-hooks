package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

type Genesis struct {
	Nonce      string `json:"nonce"`
	Timestamp  string `json:"timestamp"`
	ExtraData  string `json:"extraData"`
	GasLimit   string `json:"gasLimit"`
	GasUsed    string `json:"gasUsed"`
	Number     string `json:"number"`
	Difficulty string `json:"difficulty"`
	Coinbase   string `json:"coinbase"`
	MixHash    string `json:"mixHash"`
	ParentHash string `json:"parentHash"`

	Config Config                  `json:"config"`
	Alloc  map[string]AllocAccount `json:"alloc"`
}

type Config struct {
	ChainID             int    `json:"chainId"`
	HomesteadBlock      int    `json:"homesteadBlock"`
	EIP150Block         int    `json:"eip150Block"`
	EIP150Hash          string `json:"eip150Hash"`
	EIP155Block         int    `json:"eip155Block"`
	EIP158Block         int    `json:"eip158Block"`
	ByzantiumBlock      int    `json:"byzantiumBlock"`
	ConstantinopleBlock int    `json:"constantinopleBlock"`
	PetersburgBlock     int    `json:"petersburgBlock"`
	IstanbulBlock       int    `json:"istanbulBlock"`
	MuirGlacierBlock    int    `json:"muirglacierblock"`
	BerlinBlock         int    `json:"berlinBlock"`
	LondonBlock         int    `json:"londonBlock"`

	ContractSizeLimit int  `json:"contractSizeLimit"`
	ZeroBaseFee       bool `json:"zeroBaseFee"`

	QBFT QBFTConfig `json:"qbft"`
}

type QBFTConfig struct {
	BlockPeriodSeconds       int `json:"blockperiodseconds"`
	EpochLength              int `json:"epochlength"`
	RequestTimeoutSeconds    int `json:"requesttimeoutseconds"`
	XEmptyBlockPeriodSeconds int `json:"xemptyblockperiodseconds"`
}

type AllocAccount struct {
	Balance string `json:"balance"`
}

type UserData struct {
	Consensus             string   `json:"consensus"`
	ChainID               int      `json:"chainID"`
	BlockPeriod           int      `json:"blockperiod"`
	RequestTimeout        int      `json:"requestTimeout"`
	XEmptyBlockPeriod     bool     `json:"xemptyBlockPeriod"`
	EmptyBlockPeriod      int      `json:"emptyBlockPeriod"`
	EpochLength           int      `json:"epochLength"`
	Difficulty            int      `json:"difficulty"`
	GasLimit              string   `json:"gasLimit"`
	Coinbase              string   `json:"coinbase"`
	MaxCodeSize           int      `json:"maxCodeSize"`
	TxnSizeLimit          int      `json:"txnSizeLimit"`
	Validators            int      `json:"validators"`
	Members               int      `json:"members"`
	Bootnodes             int      `json:"bootnodes"`
	AccountPassword       string   `json:"accountPassword"`
	OutputPath            string   `json:"outputPath"`
	TesseraEnabled        bool     `json:"tesseraEnabled"`
	TesseraPassword       string   `json:"tesseraPassword"`
	QuickstartDevAccounts bool     `json:"quickstartDevAccounts"`
	NoOutputTimestamp     bool     `json:"noOutputTimestamp"`
	PrefundedAccounts     struct{} `json:"prefundedAccounts"`
	GenesisNodeAllocation string   `json:"genesisNodeAllocation"`
}

func (g *Genesis) Save(dirName string) error {
	return saveJSON(g, fmt.Sprintf("%s/besu", dirName), "genesis.json")
}

func (u *UserData) Save(dirName string) error {
	return saveJSON(u, dirName, "userData.json")
}

func saveJSON[T any](value T, dirName, fileName string) error {
	if err := os.MkdirAll(dirName, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	f, err := os.Create(fmt.Sprintf("%s/%s", dirName, fileName))
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Printf("failed to close file: %v\n", err)
		}
	}()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)

	if err := enc.Encode(value); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}
