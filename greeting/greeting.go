package greeting

import (
	"fmt"
)

// SayHello returns a string with a greeting to a given user
func SayHello(user string) string {
	return fmt.Sprintf("Hello, %s", user)
}
