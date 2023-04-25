package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/TranceAM/voting/testutil/keeper"
	"github.com/TranceAM/voting/x/voting/keeper"
	"github.com/TranceAM/voting/x/voting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.VotingKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
