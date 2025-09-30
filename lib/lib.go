package lib

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
	Consensus         string `json:"consensus"`
	ChainID           int    `json:"chainID"`
	BlockPeriod       int    `json:"blockperiod"`
	RequestTimeout    int    `json:"requestTimeout"`
	XEmptyBlockPeriod bool   `json:"xemptyBlockPeriod"`
	EmptyBlockPeriod  int    `json:"emptyBlockPeriod"`
	EpochLength       int    `json:"epochLength"`
	Difficulty        int    `json:"difficulty"`
	GasLimit          string `json:"gasLimit"`
	Coinbase          string `json:"coinbase"`
	MaxCodeSize       int    `json:"maxCodeSize"`
	TxnSizeLimit      int    `json:"txnSizeLimit"`
	Validators        int    `json:"validators"`
	AccountPassword   string `json:"accountPassword"`
	OutputPath        string `json:"outputPath"`
}
