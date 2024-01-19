package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/unigrid-project/ugdvesting-test/testutil/keeper"
	"github.com/unigrid-project/ugdvesting-test/x/ugdvesting/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.UgdvestingKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
