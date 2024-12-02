package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSupplychain{}

func NewMsgCreateSupplychain(creator string, itemCode string, quantity uint64, itemDesc string) *MsgCreateSupplychain {
	return &MsgCreateSupplychain{
		Creator:  creator,
		ItemCode: itemCode,
		Quantity: quantity,
		ItemDesc: itemDesc,
	}
}

func (msg *MsgCreateSupplychain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
