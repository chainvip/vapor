package protocol

import (
	"testing"
)

func TestNextLeaderTime(t *testing.T) {
	cases := []struct {
		desc               string
		startTime          uint64
		endTime            uint64
		now                uint64
		nodeOrder          uint64
		wantError          error
		wantNextLeaderTime int64
	}{
		{
			desc:               "normal case",
			startTime:          1557906284061,
			endTime:            1557906784061,
			now:                1557906534061,
			nodeOrder:          1,
			wantError:          nil,
			wantNextLeaderTime: 1557906537561,
		},
		{
			desc:               "best block height equals to start block height",
			startTime:          1557906284061,
			endTime:            1557906784061,
			now:                1557906284061,
			nodeOrder:          0,
			wantError:          nil,
			wantNextLeaderTime: 1557906284061,
		},
		{
			desc:               "best block height equals to start block height",
			startTime:          1557906284061,
			endTime:            1557906784061,
			now:                1557906284061,
			nodeOrder:          1,
			wantError:          nil,
			wantNextLeaderTime: 1557906284061 + blockNumEachNode*blockTimeInterval,
		},
		{
			desc:               "has no chance product block in this round of voting",
			startTime:          1557906284061,
			endTime:            1557906784061,
			now:                1557906781561,
			nodeOrder:          1,
			wantError:          errHasNoChanceProductBlock,
			wantNextLeaderTime: 0,
		},
		{
			desc:               "the node is producting block",
			startTime:          1557906284061,
			endTime:            1557906784061,
			now:                1557906284561,
			nodeOrder:          0,
			wantError:          nil,
			wantNextLeaderTime: 1557906284061,
		},
		{
			desc:               "the node is producting block",
			startTime:          1557906284061,
			endTime:            1557906784061,
			now:                1557906317561,
			nodeOrder:          1,
			wantError:          nil,
			wantNextLeaderTime: 1557906284061 + 66*blockTimeInterval,
		},
		{
			desc:               "first round, must exclude genesis block",
			startTime:          1557906284061,
			endTime:            1557906783561,
			now:                1557906286561,
			nodeOrder:          3,
			wantError:          nil,
			wantNextLeaderTime: 1557906284061 + 9*blockTimeInterval,
		},
	}

	for i, c := range cases {
		nextLeaderTime, err := nextLeaderTimeHelper(c.startTime, c.endTime, c.now, c.nodeOrder)
		if err != c.wantError {
			t.Fatalf("case #%d (%s) want error:%v, got error:%v", i, c.desc, c.wantError, err)
		}

		if err != nil {
			continue
		}
		nextLeaderTimestamp := nextLeaderTime.UnixNano() / 1e6
		if nextLeaderTimestamp != c.wantNextLeaderTime {
			t.Errorf("case #%d (%s) want next leader time:%d, got next leader time:%d", i, c.desc, c.wantNextLeaderTime, nextLeaderTimestamp)
		}
	}
}