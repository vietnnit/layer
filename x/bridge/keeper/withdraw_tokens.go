package keeper

import (
	"context"
	"encoding/hex"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tellor-io/layer/x/bridge/types"
	oracletypes "github.com/tellor-io/layer/x/oracle/types"
)

func (k Keeper) withdrawTokens(ctx context.Context, amount sdk.Coin, sender sdk.AccAddress, recipient []byte) error {
	k.Logger(ctx).Info("@withdrawTokens keeper", "amount", amount, "sender", sender, "recipient", recipient)
	// send coins from the sender to the bridge module
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(amount)); err != nil {
		return err
	}
	// burn the coins
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(amount)); err != nil {
		return err
	}

	withdrawalId, err := k.incrementWithdrawalId(ctx)
	if err != nil {
		return err
	}

	aggregate, err := k.createWithdrawalAggregate(ctx, amount, sender, recipient, withdrawalId)
	if err != nil {
		return err
	}

	k.oracleKeeper.SetAggregate(ctx, aggregate)

	// TODO: request snapshot

	return nil
}

func (k Keeper) incrementWithdrawalId(goCtx context.Context) (uint64, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	id, err := k.WithdrawalId.Get(ctx)
	if err != nil {
		id.Id = 1
		err = k.WithdrawalId.Set(ctx, id)
		if err != nil {
			return 0, err
		}
		return id.Id, nil
	}
	id.Id++
	err = k.WithdrawalId.Set(ctx, id)
	if err != nil {
		return 0, err
	}
	return id.Id, nil
}

func (k Keeper) createWithdrawalAggregate(goCtx context.Context, amount sdk.Coin, sender sdk.AccAddress, recipient []byte, withdrawalId uint64) (*oracletypes.Aggregate, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	queryId, err := k.getWithdrawalQueryId(withdrawalId)
	if err != nil {
		return nil, err
	}
	reportValue, err := k.getWithdrawalReportValue(amount, sender, recipient)
	if err != nil {
		return nil, err
	}

	aggregate := &oracletypes.Aggregate{
		QueryId:              queryId,
		AggregateValue:       hex.EncodeToString(reportValue),
		AggregateReporter:    "",
		ReporterPower:        0,
		StandardDeviation:    0,
		Reporters:            nil,
		Flagged:              false,
		Nonce:                0,
		AggregateReportIndex: 0,
		Height:               ctx.BlockHeight(),
	}
	return aggregate, nil
}

func (k Keeper) getWithdrawalQueryId(withdrawalId uint64) ([]byte, error) {
	// replicate solidity encoding,  keccak256(abi.encode(string "TRBBridge", abi.encode(uint256 withdrawalId, bool false)))

	queryTypeString := "TRBBridge"
	toLayerBool := false
	withdrawalIdUint64 := new(big.Int).SetUint64(withdrawalId)

	// prepare encoding
	StringType, err := abi.NewType("string", "", nil)
	if err != nil {
		return nil, err
	}
	Uint256Type, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return nil, err
	}
	BoolType, err := abi.NewType("bool", "", nil)
	if err != nil {
		return nil, err
	}
	BytesType, err := abi.NewType("bytes", "", nil)
	if err != nil {
		return nil, err
	}

	// encode query data arguments first
	queryDataArgs := abi.Arguments{
		{Type: BoolType},
		{Type: Uint256Type},
	}
	queryDataArgsEncoded, err := queryDataArgs.Pack(toLayerBool, withdrawalIdUint64)
	if err != nil {
		return nil, err
	}

	// encode query data
	finalArgs := abi.Arguments{
		{Type: StringType},
		{Type: BytesType},
	}
	queryDataEncoded, err := finalArgs.Pack(queryTypeString, queryDataArgsEncoded)
	if err != nil {
		return nil, err
	}

	// generate query id
	queryId := crypto.Keccak256(queryDataEncoded)

	return queryId, nil
}

func (k Keeper) getWithdrawalReportValue(amount sdk.Coin, sender sdk.AccAddress, recipient []byte) ([]byte, error) {
	// replicate solidity encoding, abi.encode(address recipient, string sender, uint256 amount)

	ethAddressString := hex.EncodeToString(recipient)
	// ensure has 0x prefix - need this?
	if ethAddressString[:2] != "0x" {
		ethAddressString = "0x" + ethAddressString
	}
	layerAddressString := sender.String()
	amountUint64 := new(big.Int).SetUint64(amount.Amount.Uint64())

	// prepare encoding
	AddressType, err := abi.NewType("address", "", nil)
	if err != nil {
		return nil, err
	}
	Uint256Type, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return nil, err
	}
	StringType, err := abi.NewType("string", "", nil)
	if err != nil {
		return nil, err
	}

	reportValueArgs := abi.Arguments{
		{Type: AddressType},
		{Type: StringType},
		{Type: Uint256Type},
	}

	// encode report value arguments
	reportValueArgsEncoded, err := reportValueArgs.Pack(ethAddressString, layerAddressString, amountUint64)
	if err != nil {
		return nil, err
	}

	return reportValueArgsEncoded, nil
}
