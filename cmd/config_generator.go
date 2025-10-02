package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/kba-tools/besu-k8s-hooks/lib"
	"github.com/urfave/cli/v3"
)

var (
	chainIDFlag = &cli.IntFlag{
		Name:  "chainID",
		Usage: "Chain ID for the blockchain",
		Value: 1337,
	}
	coinbaseFlag = &cli.StringFlag{
		Name:  "coinbase",
		Usage: "Address for earning the mining rewards",
		Value: "0x0000000000000000000000000000000000000000",
	}
	accountPasswordFlag = &cli.StringFlag{
		Name:  "accountPassword",
		Usage: "Password for the keys",
	}
	validatorsFlag = &cli.IntFlag{
		Name:  "validators",
		Usage: "Number of validators",
		Value: 4,
	}
	gasLimitFlag = &cli.StringFlag{
		Name:  "gasLimit",
		Usage: "Block gas limit (in hex)",
		Value: "0xfffffffffffffff",
	}
	blockperiodFlag = &cli.IntFlag{
		Name:  "blockperiod",
		Usage: "Time between blocks (in seconds)",
		Value: 5,
	}
	xemptyBlockPeriodFlag = &cli.BoolFlag{
		Name:  "xemptyBlockPeriod",
		Usage: "Increase the block time when there are no transactions",
	}
	emptyBlockPeriodFlag = &cli.IntFlag{
		Name:  "emptyBlockPeriod",
		Usage: "Time between blocks when there are no transactions (in seconds)",
		Value: 60,
	}
	requestTimeoutFlag = &cli.IntFlag{
		Name:  "requestTimeout",
		Usage: "Minimum request timeout for each round (in seconds)",
		Value: 30,
	}
	maxCodeSizeFlag = &cli.IntFlag{
		Name:  "maxCodeSize",
		Usage: "Maximum contract size (in KB)",
		Value: 1048576,
	}
	txnSizeLimitFlag = &cli.IntFlag{
		Name:  "txnSizeLimit",
		Usage: "Maximum transaction size (in KB)",
		Value: 1048576,
	}
	epochLengthFlag = &cli.IntFlag{
		Name:  "epochLength",
		Usage: "Number of blocks after which the votes are reset",
		Value: 30000,
	}
	difficultyFlag = &cli.IntFlag{
		Name:  "difficulty",
		Usage: "Difficulty of the network",
		Value: 0x1,
	}
	outputFlag = &cli.StringFlag{
		Name:  "output",
		Usage: "Output file location",
	}
)

var app = newApp("Besu Config Generator")

func newApp(usage string) *cli.Command {
	app := &cli.Command{}
	app.EnableShellCompletion = true
	app.Usage = usage
	app.Copyright = "Copyright 2025 Kerala Blockchain Academy"
	return app
}

func init() {
	app.Name = "besu-config-generator"
	app.Flags = []cli.Flag{
		chainIDFlag,
		coinbaseFlag,
		accountPasswordFlag,
		validatorsFlag,
		gasLimitFlag,
		blockperiodFlag,
		xemptyBlockPeriodFlag,
		emptyBlockPeriodFlag,
		requestTimeoutFlag,
		maxCodeSizeFlag,
		txnSizeLimitFlag,
		epochLengthFlag,
		difficultyFlag,
		outputFlag,
	}
	app.Action = generate
}

func generate(ctx context.Context, c *cli.Command) error {
	genesis := &lib.Genesis{
		Timestamp:  hexutil.EncodeUint64(uint64(time.Now().Unix())),
		Coinbase:   c.String(coinbaseFlag.Name),
		Nonce:      "0x0",
		GasLimit:   c.String(gasLimitFlag.Name),
		GasUsed:    "0x0",
		Number:     "0x0",
		Difficulty: c.String(difficultyFlag.Name),
		ParentHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
		MixHash:    "0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365",
		ExtraData:  "0xf87aa00000000000000000000000000000000000000000000000000000000000000000f854948001d7a021b97ba1ee3db980e7b9951464e3d08494859e38d3039470de5edb161f205d74b80d7cf9b0949e33e1d20ccfcc4976b509d4ff852e2b8efcd504943333630c724d5d6dd53159c577114a1f0ce1c02bc080c0",
		Config: lib.Config{
			ChainID:             c.Int(chainIDFlag.Name),
			HomesteadBlock:      0,
			EIP150Block:         0,
			EIP150Hash:          "0x0000000000000000000000000000000000000000000000000000000000000000",
			EIP155Block:         0,
			EIP158Block:         0,
			ByzantiumBlock:      0,
			ConstantinopleBlock: 0,
			PetersburgBlock:     0,
			IstanbulBlock:       0,
			MuirGlacierBlock:    0,
			BerlinBlock:         0,
			LondonBlock:         0,
			ZeroBaseFee:         true,
			ContractSizeLimit:   c.Int(maxCodeSizeFlag.Name),
			QBFT: lib.QBFTConfig{
				BlockPeriodSeconds:       c.Int(blockperiodFlag.Name),
				XEmptyBlockPeriodSeconds: c.Int(emptyBlockPeriodFlag.Name),
				RequestTimeoutSeconds:    c.Int(requestTimeoutFlag.Name),
				EpochLength:              c.Int(epochLengthFlag.Name),
			},
		},
	}

	userData := &lib.UserData{
		Consensus:         "qbft",
		ChainID:           c.Int(chainIDFlag.Name),
		Coinbase:          c.String(coinbaseFlag.Name),
		BlockPeriod:       c.Int(blockperiodFlag.Name),
		RequestTimeout:    c.Int(requestTimeoutFlag.Name),
		XEmptyBlockPeriod: c.Bool(xemptyBlockPeriodFlag.Name),
		EmptyBlockPeriod:  c.Int(emptyBlockPeriodFlag.Name),
		EpochLength:       c.Int(epochLengthFlag.Name),
		Difficulty:        c.Int(difficultyFlag.Name),
		GasLimit:          c.String(gasLimitFlag.Name),
		MaxCodeSize:       c.Int(maxCodeSizeFlag.Name),
		TxnSizeLimit:      c.Int(txnSizeLimitFlag.Name),
		Validators:        c.Int(validatorsFlag.Name),
		AccountPassword:   c.String(accountPasswordFlag.Name),
		OutputPath:        c.String(outputFlag.Name),
	}

	if err := genesis.Save(c.String(outputFlag.Name)); err != nil {
		return err
	}

	if err := userData.Save(c.String(outputFlag.Name)); err != nil {
		return err
	}

	if err := lib.SaveConfigTOML(c.String(outputFlag.Name), c.Int(chainIDFlag.Name), "0.0.0.0", 30303); err != nil {
		return err
	}

	for i := range c.Int(validatorsFlag.Name) {
		dirName := fmt.Sprintf("%s/validator%d", c.String(outputFlag.Name), i)
		if err := os.MkdirAll(dirName, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		if err := lib.GenerateValidatorKeys(dirName, c.String(accountPasswordFlag.Name)); err != nil {
			return err
		}

	}

	return nil
}

func main() {
	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
