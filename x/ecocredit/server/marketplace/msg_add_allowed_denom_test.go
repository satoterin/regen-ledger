package marketplace

import (
	"encoding/json"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

type addAllowedDenomSuite struct {
	*baseSuite
	err error
}

func TestAddAllowedDenom(t *testing.T) {
	gocuke.NewRunner(t, &addAllowedDenomSuite{}).Path("./features/msg_add_allowed_denom.feature").Run()
}

func (s *addAllowedDenomSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t, 1)
}

func (s *addAllowedDenomSuite) AnAllowedDenomWithProperties(a gocuke.DocString) {
	var msg *marketplace.MsgAddAllowedDenom

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	err = s.k.stateStore.AllowedDenomTable().Insert(s.ctx, &api.AllowedDenom{
		BankDenom:    msg.BankDenom,
		DisplayDenom: msg.DisplayDenom,
		Exponent:     msg.Exponent,
	})
	require.NoError(s.t, err)
}

func (s *addAllowedDenomSuite) AliceAttemptsToAddADenomWithProperties(a gocuke.DocString) {
	var msg *marketplace.MsgAddAllowedDenom

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	_, s.err = s.k.AddAllowedDenom(s.ctx, msg)
}

func (s *addAllowedDenomSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *addAllowedDenomSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *addAllowedDenomSuite) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}
