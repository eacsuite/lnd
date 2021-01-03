// +build rpctest

package aezeed

import "github.com/eacsuite/eacwallet/waddrmgr"

func init() {
	// For the purposes of our itest, we'll crank down the scrypt params a
	// bit.
	scryptN = waddrmgr.FastScryptOptions.N
	scryptR = waddrmgr.FastScryptOptions.R
	scryptP = waddrmgr.FastScryptOptions.P
}
