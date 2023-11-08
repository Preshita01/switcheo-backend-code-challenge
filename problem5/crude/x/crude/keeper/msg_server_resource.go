package keeper

import (
	"context"
	"fmt"

	"crude/x/crude/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateResource(goCtx context.Context, msg *types.MsgCreateResource) (*types.MsgCreateResourceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var resource = types.Resource{
		Creator:   msg.Creator,
		RName:     msg.RName,
		RCategory: msg.RCategory,
		RColour:   msg.RColour,
		RSize:     msg.RSize,
		RQuantity: msg.RQuantity,
	}

	id := k.AppendResource(
		ctx,
		resource,
	)

	return &types.MsgCreateResourceResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateResource(goCtx context.Context, msg *types.MsgUpdateResource) (*types.MsgUpdateResourceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var resource = types.Resource{
		Creator:   msg.Creator,
		Id:        msg.Id,
		RName:     msg.RName,
		RCategory: msg.RCategory,
		RColour:   msg.RColour,
		RSize:     msg.RSize,
		RQuantity: msg.RQuantity,
	}

	// Checks that the element exists
	val, found := k.GetResource(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetResource(ctx, resource)

	return &types.MsgUpdateResourceResponse{}, nil
}

func (k msgServer) DeleteResource(goCtx context.Context, msg *types.MsgDeleteResource) (*types.MsgDeleteResourceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetResource(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveResource(ctx, msg.Id)

	return &types.MsgDeleteResourceResponse{}, nil
}
