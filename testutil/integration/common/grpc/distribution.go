// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)
package grpc

import (
	"context"

	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

// GetDelegationTotalRewards returns the total delegation rewards for the given delegator.
func (gqh *IntegrationHandler) GetDelegationTotalRewards(delegatorAddress string) (*distrtypes.QueryDelegationTotalRewardsResponse, error) {
	distrClient := gqh.network.GetDistrClient()
	return distrClient.DelegationTotalRewards(context.Background(), &distrtypes.QueryDelegationTotalRewardsRequest{
		DelegatorAddress: delegatorAddress,
	})
}

// GetDelegationTotalRewards returns the total delegation rewards for the given delegator.
func (gqh *IntegrationHandler) GetDelegatorWithdrawAddr(delegatorAddress string) (*distrtypes.QueryDelegatorWithdrawAddressResponse, error) {
	distrClient := gqh.network.GetDistrClient()
	return distrClient.DelegatorWithdrawAddress(context.Background(), &distrtypes.QueryDelegatorWithdrawAddressRequest{
		DelegatorAddress: delegatorAddress,
	})
}

// GetDelegationTotalRewards returns the total delegation rewards for the given delegator.
func (gqh *IntegrationHandler) GetValidatorCommission(validatorAddress string) (*distrtypes.QueryValidatorCommissionResponse, error) {
	distrClient := gqh.network.GetDistrClient()
	return distrClient.ValidatorCommission(context.Background(), &distrtypes.QueryValidatorCommissionRequest{
		ValidatorAddress: validatorAddress,
	})
}