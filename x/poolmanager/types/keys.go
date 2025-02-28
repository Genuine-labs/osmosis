package types

import (
	"fmt"
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
)

const (
	ModuleName   = "poolmanager"
	KeySeparator = "|"

	StoreKey = ModuleName

	RouterKey = ModuleName
)

var (
	// KeyNextGlobalPoolId defines key to store the next Pool ID to be used.
	KeyNextGlobalPoolId = []byte{0x01}

	// SwapModuleRouterPrefix defines prefix to store pool id to swap module mappings.
	SwapModuleRouterPrefix = []byte{0x02}

	// KeyPoolVolumePrefix defines prefix to store pool volume.
	KeyPoolVolumePrefix = []byte{0x03}

	// DenomTradePairPrefix defines prefix to store denom trade pair for taker fee.
	DenomTradePairPrefix = []byte{0x04}

	// KeyTakerFeeStakersProtoRev defines key to store the taker fee for stakers tracker.
	KeyTakerFeeStakersProtoRev = []byte{0x05}

	// KeyTakerFeeCommunityPoolProtoRev defines key to store the taker fee for community pool tracker.
	KeyTakerFeeCommunityPoolProtoRev = []byte{0x06}

	// KeyTakerFeeProtoRevAccountingHeight defines key to store the accounting height for the above taker fee trackers.
	KeyTakerFeeProtoRevAccountingHeight = []byte{0x07}
)

// ModuleRouteToBytes serializes moduleRoute to bytes.
func FormatModuleRouteKey(poolId uint64) []byte {
	return []byte(fmt.Sprintf("%s%d", SwapModuleRouterPrefix, poolId))
}

// FormatDenomTradePairKey serializes denom trade pair to bytes.
// Denom trade pair is automatically sorted lexicographically.
func FormatDenomTradePairKey(denom0, denom1 string) []byte {
	denoms := []string{denom0, denom1}
	sort.Strings(denoms)
	return []byte(fmt.Sprintf("%s%s%s%s%s", DenomTradePairPrefix, KeySeparator, denoms[0], KeySeparator, denoms[1]))
}

// ParseModuleRouteFromBz parses the raw bytes into ModuleRoute.
// Returns error if fails to parse or if the bytes are empty.
func ParseModuleRouteFromBz(bz []byte) (ModuleRoute, error) {
	moduleRoute := ModuleRoute{}
	err := proto.Unmarshal(bz, &moduleRoute)
	if err != nil {
		return ModuleRoute{}, err
	}
	return moduleRoute, err
}

// KeyPoolVolume returns the key for the pool volume corresponding to the given poolId.
func KeyPoolVolume(poolId uint64) []byte {
	return []byte(fmt.Sprintf("%s%s%d%s", KeyPoolVolumePrefix, KeySeparator, poolId, KeySeparator))
}

// ParseDenomTradePairKey parses the raw bytes of the DenomTradePairKey into a denom trade pair.
func ParseDenomTradePairKey(key []byte) (denom0, denom1 string, err error) {
	keyStr := string(key)
	parts := strings.Split(keyStr, KeySeparator)

	denom0 = parts[1]
	denom1 = parts[2]

	err = sdk.ValidateDenom(denom0)
	if err != nil {
		return "", "", err
	}

	err = sdk.ValidateDenom(denom1)
	if err != nil {
		return "", "", err
	}

	return denom0, denom1, nil
}
