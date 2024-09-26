package sp5_test

import (
	"testing"

	keepertest "github.com/cel031/SP5/testutil/keeper"
	"github.com/cel031/SP5/testutil/nullify"
	"github.com/cel031/SP5/x/sp5"
	"github.com/cel031/SP5/x/sp5/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Sp5Keeper(t)
	sp5.InitGenesis(ctx, *k, genesisState)
	got := sp5.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}