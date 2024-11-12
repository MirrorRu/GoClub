package errapp

import "fmt"

// Wrap wraps error with given message
func Wrap(topMsg string, wrappedErr error) error {
	if wrappedErr == nil {
		return nil
	}
	return fmt.Errorf(topMsg+": %w", wrappedErr)

}
