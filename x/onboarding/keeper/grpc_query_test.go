package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/QOM-One/QomApp/v7/x/onboarding/types"
)

func (suite *KeeperTestSuite) TestQueryParams() {
	ctx := sdk.WrapSDKContext(suite.ctx)
	expParams := types.DefaultParams()

	res, err := suite.queryClient.Params(ctx, &types.QueryParamsRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(expParams, res.Params)
}
