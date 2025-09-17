//go:build !with_antithesis_sdk

package antithesis

// Always asserts that condition is true every time this function is called, and that it is called at least once.
func Always(condition bool, message string, details map[string]any) {
	// no-op when Antithesis SDK is not enabled
}

// AlwaysOrUnreachable asserts that condition is true every time this function is called.
func AlwaysOrUnreachable(condition bool, message string, details map[string]any) {
	// no-op when Antithesis SDK is not enabled
}

// Sometimes asserts that condition is true at least one time that this function was called.
func Sometimes(condition bool, message string, details map[string]any) {
	// no-op when Antithesis SDK is not enabled
}

// Unreachable asserts that a line of code is never reached.
func Unreachable(message string, details map[string]any) {
	// no-op when Antithesis SDK is not enabled
}

// Reachable asserts that a line of code is reached at least once.
func Reachable(message string, details map[string]any) {
	// no-op when Antithesis SDK is not enabled
}
