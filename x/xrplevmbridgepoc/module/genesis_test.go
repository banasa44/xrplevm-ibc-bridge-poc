package xrplevmbridgepoc_test

import (
	"testing"

	keepertest "xrplevm-bridge-poc/testutil/keeper"
	"xrplevm-bridge-poc/testutil/nullify"
	xrplevmbridgepoc "xrplevm-bridge-poc/x/xrplevmbridgepoc/module"
	"xrplevm-bridge-poc/x/xrplevmbridgepoc/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.XrplevmbridgepocKeeper(t)
	xrplevmbridgepoc.InitGenesis(ctx, k, genesisState)
	got := xrplevmbridgepoc.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
