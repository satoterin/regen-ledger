package basket

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func (k Keeper) BasketBalances(ctx context.Context, request *baskettypes.QueryBasketBalancesRequest) (*baskettypes.QueryBasketBalancesResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	basket, err := k.stateStore.BasketTable().GetByBasketDenom(ctx, request.BasketDenom)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return nil, sdkerrors.Wrapf(err, "basket %s not found", request.BasketDenom)
		}
		return nil, sdkerrors.Wrapf(err, "failed to get basket %s", request.BasketDenom)
	}

	pulsarPageReq, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.BasketBalanceTable().List(ctx, api.BasketBalancePrimaryKey{}.WithBasketId(basket.Id),
		ormlist.Paginate(pulsarPageReq),
	)
	if err != nil {
		return nil, err
	}

	res := &baskettypes.QueryBasketBalancesResponse{}
	for it.Next() {
		bal, err := it.Value()
		if err != nil {
			return nil, err
		}

		balanceGogo := &baskettypes.BasketBalance{}
		if err = ormutil.PulsarToGogoSlow(bal, balanceGogo); err != nil {
			return nil, err
		}
		res.Balances = append(res.Balances, balanceGogo)

		res.BalancesInfo = append(res.BalancesInfo, &baskettypes.BasketBalanceInfo{
			BatchDenom: bal.BatchDenom,
			Balance:    bal.Balance,
		})
	}
	it.Close()

	res.Pagination, err = ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	return res, err
}
