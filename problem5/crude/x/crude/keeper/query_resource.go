package keeper

import (
	"context"

	"crude/x/crude/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ResourceAll(goCtx context.Context, req *types.QueryAllResourceRequest) (*types.QueryAllResourceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var resources []types.Resource
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	resourceStore := prefix.NewStore(store, types.KeyPrefix(types.ResourceKey))

	pageRes, err := query.Paginate(resourceStore, req.Pagination, func(key []byte, value []byte) error {
		var resource types.Resource
		if err := k.cdc.Unmarshal(value, &resource); err != nil {
			return err
		}

		resources = append(resources, resource)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllResourceResponse{Resource: resources, Pagination: pageRes}, nil
}

func (k Keeper) Resource(goCtx context.Context, req *types.QueryGetResourceRequest) (*types.QueryGetResourceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	resource, found := k.GetResource(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetResourceResponse{Resource: resource}, nil
}
