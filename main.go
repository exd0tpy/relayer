package main

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/exd0tpy/relayer/cmd"
)

func main() {
	cmd.Execute()
}

func init() {
	//prevent incorrect bech32 address prefixed addresses when calling AccAddress.String()
	sdk.SetAddrCacheEnabled(false)
}
