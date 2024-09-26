package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cel031/SP5/testutil/keeper"
	"github.com/cel031/SP5/x/timtam/keeper"
	"github.com/cel031/SP5/x/timtam/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TimtamKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
