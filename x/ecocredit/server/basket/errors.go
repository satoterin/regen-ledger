package basket

import sdkerrors "github.com/cosmos/cosmos-sdk/errors"

var codespace = "ecocredit/basket"

var (
	ErrCantDisableRetire = sdkerrors.Register(codespace, 2, "can't disable retirement when taking from this basket")
)
