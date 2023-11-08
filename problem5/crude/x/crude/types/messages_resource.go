package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateResource = "create_resource"
	TypeMsgUpdateResource = "update_resource"
	TypeMsgDeleteResource = "delete_resource"
)

var _ sdk.Msg = &MsgCreateResource{}

func NewMsgCreateResource(creator string, rName string, rCategory string, rColour string, rSize string, rQuantity uint64) *MsgCreateResource {
	return &MsgCreateResource{
		Creator:   creator,
		RName:     rName,
		RCategory: rCategory,
		RColour:   rColour,
		RSize:     rSize,
		RQuantity: rQuantity,
	}
}

func (msg *MsgCreateResource) Route() string {
	return RouterKey
}

func (msg *MsgCreateResource) Type() string {
	return TypeMsgCreateResource
}

func (msg *MsgCreateResource) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateResource) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateResource) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateResource{}

func NewMsgUpdateResource(creator string, id uint64, rName string, rCategory string, rColour string, rSize string, rQuantity uint64) *MsgUpdateResource {
	return &MsgUpdateResource{
		Id:        id,
		Creator:   creator,
		RName:     rName,
		RCategory: rCategory,
		RColour:   rColour,
		RSize:     rSize,
		RQuantity: rQuantity,
	}
}

func (msg *MsgUpdateResource) Route() string {
	return RouterKey
}

func (msg *MsgUpdateResource) Type() string {
	return TypeMsgUpdateResource
}

func (msg *MsgUpdateResource) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateResource) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateResource) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteResource{}

func NewMsgDeleteResource(creator string, id uint64) *MsgDeleteResource {
	return &MsgDeleteResource{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteResource) Route() string {
	return RouterKey
}

func (msg *MsgDeleteResource) Type() string {
	return TypeMsgDeleteResource
}

func (msg *MsgDeleteResource) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteResource) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteResource) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
