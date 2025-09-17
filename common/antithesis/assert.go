//go:build with_antithesis_sdk

package antithesis

//nolint:depguard | Use of Antithesis SDK is allowed (only) here
import "github.com/antithesishq/antithesis-sdk-go/assert"

var (
	Always              = assert.Always
	AlwaysOrUnreachable = assert.AlwaysOrUnreachable
	Sometimes           = assert.Sometimes
	Unreachable         = assert.Unreachable
	Reachable           = assert.Reachable
)
