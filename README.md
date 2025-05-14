用于实现离线机器绑定授权机制：包括指纹生成、授权文件签名生成、运行时校验逻辑。该方案采用 RSA 签名，适合离线部署场景

## 项目结构
```aiignore
machine-auth/
├── cmd/
│   ├── gen-fingerprint.go     # 客户端运行，生成指纹
│   ├── gen-license.go         # 开发者使用，生成授权文件
│   └── verify.go              # 客户端启动时调用，验证授权
├── internal/
│   ├── fingerprint/
│   │   └── fingerprint.go     # 硬件信息采集逻辑
│   ├── license/
│   │   └── license.go         # 授权文件签名/验证
├── keys/
│   ├── private.pem            # 私钥（仅开发者持有）
│   └── public.pem             # 公钥（嵌入客户端）
```

## RSA 密钥生成
```aiignore
# 生成 RSA 密钥对
openssl genrsa -out private.pem 2048
openssl rsa -in private.pem -pubout -out public.pem
```