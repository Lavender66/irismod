package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/irisnet/irismod/modules/rental/types"
)

// Rent set or update rental info for an nft.
func (k Keeper) Rent(ctx sdk.Context, rental types.RentalInfo) error {
	// this nft must expire if to be set again.
	// FIXME: proto should use int64 or Time than uint64
	rental, exist := k.GetRentalInfo(ctx, rental.ClassId, rental.NftId)
	if exist && ctx.BlockTime().Unix() < int64(rental.Expires) {
		return sdkerrors.Wrapf(types.ErrNotArriveExpires, "Expires is (%d)", rental.Expires)
	}

	// set rental info
	k.setRentalInfo(ctx, rental.ClassId, rental.NftId, rental.User, rental.Expires)
	return nil
}

// setRentalInfo sets the rental info for an nft.
func (k Keeper) setRentalInfo(ctx sdk.Context,
	classId, nftId, user string,
	expires uint64) {
	store := ctx.KVStore(k.storeKey)
	r := types.RentalInfo{
		User:    user,
		ClassId: classId,
		NftId:   nftId,
		Expires: expires,
	}
	bz := k.cdc.MustMarshal(&r)
	store.Set(rentalInfoKey(r.ClassId, r.NftId), bz)
}

// GetRentalInfo returns the rental info for an nft.
func (k Keeper) GetRentalInfo(ctx sdk.Context,
	classId, nftId string) (types.RentalInfo, bool) {
	var v types.RentalInfo
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(rentalInfoKey(classId, nftId))
	if bz == nil {
		return types.RentalInfo{}, false
	}
	k.cdc.MustUnmarshal(bz, &v)
	return v, true
}

// GetRentalInfos returns all rental infos.
func (k Keeper) GetRentalInfos(ctx sdk.Context) (ris []types.RentalInfo) {
	store := ctx.KVStore(k.storeKey)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var rental types.RentalInfo
		k.cdc.MustUnmarshal(iterator.Value(), &rental)
		ris = append(ris, rental)
	}
	return ris
}