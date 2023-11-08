package keeper

import (
	"encoding/binary"

	"crude/x/crude/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetResourceCount get the total number of resource
func (k Keeper) GetResourceCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ResourceCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetResourceCount set the total number of resource
func (k Keeper) SetResourceCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ResourceCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendResource appends a resource in the store with a new id and update the count
func (k Keeper) AppendResource(
	ctx sdk.Context,
	resource types.Resource,
) uint64 {
	// Create the resource
	count := k.GetResourceCount(ctx)

	// Set the ID of the appended value
	resource.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResourceKey))
	appendedValue := k.cdc.MustMarshal(&resource)
	store.Set(GetResourceIDBytes(resource.Id), appendedValue)

	// Update resource count
	k.SetResourceCount(ctx, count+1)

	return count
}

// SetResource set a specific resource in the store
func (k Keeper) SetResource(ctx sdk.Context, resource types.Resource) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResourceKey))
	b := k.cdc.MustMarshal(&resource)
	store.Set(GetResourceIDBytes(resource.Id), b)
}

// GetResource returns a resource from its id
func (k Keeper) GetResource(ctx sdk.Context, id uint64) (val types.Resource, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResourceKey))
	b := store.Get(GetResourceIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveResource removes a resource from the store
func (k Keeper) RemoveResource(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResourceKey))
	store.Delete(GetResourceIDBytes(id))
}

// GetAllResource returns all resource
func (k Keeper) GetAllResource(ctx sdk.Context) (list []types.Resource) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ResourceKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Resource
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetResourceIDBytes returns the byte representation of the ID
func GetResourceIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetResourceIDFromBytes returns ID in uint64 format from a byte array
func GetResourceIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
