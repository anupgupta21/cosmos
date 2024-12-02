package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/username/helloapp/testutil/keeper"
	"github.com/username/helloapp/x/helloapp/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.HelloappKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
