package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cel031/SP5/testutil/keeper"
	"github.com/cel031/SP5/x/sp5/keeper"
	"github.com/cel031/SP5/x/sp5/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.Sp5Keeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
