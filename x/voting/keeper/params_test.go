package keeper_test

import (
	"testing"

	testkeeper "github.com/TranceAM/voting/testutil/keeper"
	"github.com/TranceAM/voting/x/voting/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.VotingKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
