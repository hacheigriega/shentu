package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params/subspace"
	params "github.com/cosmos/cosmos-sdk/x/params/subspace"
)

var (
	ParamsStoreKeyTaskParams = []byte("taskparams")
	ParamsStoreKeyPoolParams = []byte("poolparams")
)

// Default parameters
var (
	DefaultExpirationDuration = time.Duration(24) * time.Hour
	DefaultAggregationWindow  = int64(20)
	DefaultAggregationResult  = sdk.NewInt(1)
	DefaultEpsilon1           = sdk.NewInt(1)
	DefaultEpsilon2           = sdk.NewInt(100)
	DefaultThresholdScore     = sdk.NewInt(128)

	DefaultLockedInBlocks    = int64(30)
	DefaultMinimumCollateral = int64(50000)
)

// ParamKeyTable is the key declaration for parameters.
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		params.NewParamSetPair(ParamsStoreKeyTaskParams, TaskParams{}, validateTaskParams),
		params.NewParamSetPair(ParamsStoreKeyPoolParams, LockedPoolParams{}, validatePoolParams),
	)
}

type TaskParams struct {
	ExpirationDuration time.Duration `json:"task_expiration_duration"`
	AggregationWindow  int64         `json:"task_aggregation_window"`
	AggregationResult  sdk.Int       `json:"task_aggregation_result"`
	ThresholdScore     sdk.Int       `json:"task_threshold_score"`
	Epsilon1           sdk.Int       `json:"task_epsilon1"`
	Epsilon2           sdk.Int       `json:"task_epsilon2"`
}

// NewTaskParams returns a TaskParams object.
func NewTaskParams(expirationDuration time.Duration, aggregationWindow int64, aggregationResult,
	thredholdScore, epsilon1, epsilon2 sdk.Int) TaskParams {
	return TaskParams{
		ExpirationDuration: expirationDuration,
		AggregationWindow:  aggregationWindow,
		AggregationResult:  aggregationResult,
		ThresholdScore:     thredholdScore,
		Epsilon1:           epsilon1,
		Epsilon2:           epsilon2,
	}
}

// DefaultTaskParams generates default set for TaskParams.
func DefaultTaskParams() TaskParams {
	return NewTaskParams(DefaultExpirationDuration, DefaultAggregationWindow,
		DefaultAggregationResult, DefaultThresholdScore, DefaultEpsilon1, DefaultEpsilon2)
}

func validateTaskParams(i interface{}) error {
	taskParams, ok := i.(TaskParams)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if taskParams.ExpirationDuration < 0 ||
		taskParams.AggregationWindow < 0 ||
		taskParams.ThresholdScore.GT(sdk.NewInt(255)) ||
		taskParams.Epsilon1.LT(sdk.NewInt(0)) ||
		taskParams.Epsilon2.LT(sdk.NewInt(0)) {
		return ErrInvalidTaskParams
	}
	return nil
}

type LockedPoolParams struct {
	LockedInBlocks    int64 `json:"locked_in_blocks"`
	MinimumCollateral int64 `json:"minimum_collateral"`
}

// NewLockedPoolParams returns a LockedPoolParams object.
func NewLockedPoolParams(lockedInBlocks, minimumCollateral int64) LockedPoolParams {
	return LockedPoolParams{
		LockedInBlocks:    lockedInBlocks,
		MinimumCollateral: minimumCollateral,
	}
}

// DefaultLockedPoolParams generates default set for LockedPoolParams
func DefaultLockedPoolParams() LockedPoolParams {
	return NewLockedPoolParams(DefaultLockedInBlocks, DefaultMinimumCollateral)
}

func validatePoolParams(i interface{}) error {
	poolParams, ok := i.(LockedPoolParams)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if poolParams.LockedInBlocks < 0 {
		return ErrInvalidPoolParams
	}
	if poolParams.MinimumCollateral < 0 {
		return ErrInvalidPoolParams
	}
	return nil
}

type ParamSubspace interface {
	Get(ctx sdk.Context, key []byte, ptr interface{})
	Set(ctx sdk.Context, key []byte, param interface{})
	WithKeyTable(table subspace.KeyTable) subspace.Subspace
}
