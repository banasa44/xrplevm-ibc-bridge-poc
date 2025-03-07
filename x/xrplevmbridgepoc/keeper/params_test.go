package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "xrplevm-bridge-poc/testutil/keeper"
	"xrplevm-bridge-poc/x/xrplevmbridgepoc/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.XrplevmbridgepocKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
