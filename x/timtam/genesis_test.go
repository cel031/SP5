package timtam_test

import (
	"testing"

	keepertest "github.com/cel031/SP5/testutil/keeper"
	"github.com/cel031/SP5/testutil/nullify"
	"github.com/cel031/SP5/x/timtam"
	"github.com/cel031/SP5/x/timtam/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TimtamKeeper(t)
	timtam.InitGenesis(ctx, *k, genesisState)
	got := timtam.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
