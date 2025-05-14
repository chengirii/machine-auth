// cmd/gen-fingerprint.go
package main

import (
	"fmt"
	"machine-auth/internal/fingerprint"
)

func main() {
	fmt.Println(fingerprint.GetMachineFingerprint())
}
