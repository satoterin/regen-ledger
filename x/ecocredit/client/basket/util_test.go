package basketclient

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil"

	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func TestParseBasketCredits(t *testing.T) {
	emptyJson := testutil.WriteToNewTempFile(t, `{}`).Name()
	invalidJson := testutil.WriteToNewTempFile(t, `{foo:bar}`).Name()
	duplicateJson := testutil.WriteToNewTempFile(t, `{"foo":"bar","foo":"baz"}`).Name()
	validJson := testutil.WriteToNewTempFile(t, `[
		{
			"batch_denom": "C01-001-20210101-20210101-001",
			"amount": "10"
		},
		{
			"batch_denom": "C01-001-20210101-20210101-002",
			"amount": "2.5"
		}
	]`).Name()

	testCases := []struct {
		name      string
		file      string
		expErr    bool
		expErrMsg string
		expRes    []*basket.BasketCredit
	}{
		{
			name:      "empty file path",
			file:      "",
			expErr:    true,
			expErrMsg: "no such file or directory",
		},
		{
			name:      "empty json object",
			file:      emptyJson,
			expErr:    true,
			expErrMsg: "cannot unmarshal object",
		},
		{
			name:      "invalid json format",
			file:      invalidJson,
			expErr:    true,
			expErrMsg: "invalid character",
		},
		{
			name:      "duplicate json keys",
			file:      duplicateJson,
			expErr:    true,
			expErrMsg: "duplicate key",
		},
		{
			name: "valid test",
			file: validJson,
			expRes: []*basket.BasketCredit{
				{
					BatchDenom: "C01-001-20210101-20210101-001",
					Amount:     "10",
				},
				{
					BatchDenom: "C01-001-20210101-20210101-002",
					Amount:     "2.5",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := parseBasketCredits(tc.file)
			if tc.expErr {
				require.Error(t, err)
				require.ErrorContains(t, err, tc.expErrMsg)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expRes, res)
			}
		})
	}
}
