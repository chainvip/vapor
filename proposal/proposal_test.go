package proposal

import (
	"testing"

	"github.com/vapor/consensus"
)

func TestCreateCoinbaseTx(t *testing.T) {
	consensus.ActiveNetParams = consensus.Params{
		ProducerSubsidys: []consensus.ProducerSubsidy{
			{BeginBlock: 0, EndBlock: 0, Subsidy: 24},
			{BeginBlock: 1, EndBlock: 840000, Subsidy: 24},
			{BeginBlock: 840001, EndBlock: 1680000, Subsidy: 12},
			{BeginBlock: 1680001, EndBlock: 3360000, Subsidy: 6},
		},
	}
	reductionInterval := uint64(840000)
	cases := []struct {
		height  uint64
		txFee   uint64
		subsidy uint64
	}{
		{
			height:  reductionInterval - 1,
			txFee:   100000000,
			subsidy: 24 + 100000000,
		},
		{
			height:  reductionInterval,
			txFee:   2000000000,
			subsidy: 24 + 2000000000,
		},
		{
			height:  reductionInterval + 1,
			txFee:   0,
			subsidy: 12,
		},
		{
			height:  reductionInterval * 2,
			txFee:   100000000,
			subsidy: 12 + 100000000,
		},
	}

	for index, c := range cases {
		coinbaseTx, err := createCoinbaseTx(nil, c.txFee, c.height)
		if err != nil {
			t.Fatal(err)
		}

		outputAmount := coinbaseTx.Outputs[0].OutputCommitment().Amount
		if outputAmount != c.subsidy {
			t.Fatalf("index:%d,coinbase tx reward dismatch, expected: %d, have: %d", index, c.subsidy, outputAmount)
		}
	}
}