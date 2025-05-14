// cmd/gen-license.go
package main

import (
	"encoding/json"
	"fmt"
	"machine-auth/internal/license"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("用法: gen-license <fingerprint>")
		return
	}

	// 获取指纹
	fp := os.Args[1]

	// 加载私钥
	privKey, err := license.LoadPrivateKey("keys/private.pem")
	if err != nil {
		fmt.Println("无法加载私钥:", err)
		return
	}

	// 创建授权数据
	lic := &license.License{
		Fingerprint: fp,
		ValidUntil:  time.Date(9999, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	// 使用私钥签名授权文件
	err = license.SignLicense(lic, privKey)
	if err != nil {
		fmt.Println("无法签名授权文件:", err)
		return
	}

	// 将授权文件序列化并写入磁盘
	data, err := json.MarshalIndent(lic, "", "  ")
	if err != nil {
		fmt.Println("序列化授权文件失败:", err)
		return
	}

	err = os.WriteFile("license.sig", data, 0644)
	if err != nil {
		fmt.Println("写入授权文件失败:", err)
		return
	}

	// 成功提示
	fmt.Println("已生成授权文件 license.sig")
}
