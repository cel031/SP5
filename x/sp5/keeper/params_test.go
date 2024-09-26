package keeper_test

import (
	"testing"

	testkeeper "github.com/cel031/SP5/testutil/keeper"
	"github.com/cel031/SP5/x/sp5/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Sp5Keeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
