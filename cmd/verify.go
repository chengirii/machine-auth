// cmd/verify.go
package main

import (
	"encoding/json"
	"fmt"
	"machine-auth/internal/fingerprint"
	"machine-auth/internal/license"
	"os"
	"time"
)

func main() {
	pubKey, _ := license.LoadPublicKey("keys/public.pem")
	data, _ := os.ReadFile("license.sig")
	var lic license.License
	_ = json.Unmarshal(data, &lic)

	fp := fingerprint.GetMachineFingerprint()
	if fp != lic.Fingerprint {
		fmt.Println(" 硬件不匹配，禁止启动")
		os.Exit(1)
	}
	if err := license.VerifyLicense(&lic, pubKey); err != nil {
		fmt.Println("授权验证失败：", err)
		os.Exit(1)
	}
	if lic.ValidUntil.Before(time.Now()) {
		fmt.Println(" 授权已过期")
		os.Exit(1)
	}

	fmt.Println("授权通过，允许运行")
}
