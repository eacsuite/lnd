package lnwallet

import (
	"github.com/eacsuite/eacutil"
	"github.com/eacsuite/eacwallet/wallet/txrules"
	"github.com/eacsuite/lnd/input"
)

// DefaultDustLimit is used to calculate the dust HTLC amount which will be
// send to other node during funding process.
func DefaultDustLimit() eacutil.Amount {
	return txrules.GetDustThreshold(input.P2WSHSize, txrules.DefaultRelayFeePerKb)
}
