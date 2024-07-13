package util

import (
	"crypto/sha1"
	"fmt"
	"strconv"
	"strings"
)

func GenerateSHA1Signature(time int64, appID, secret string) (string, error) {
	// 将输入串接
	input := strings.Join([]string{appID, strconv.FormatInt(time, 10), secret}, "&")

	// 创建一个新的SHA-1哈希对象
	hasher := sha1.New()

	// 写入要哈希的数据
	if _, err := hasher.Write([]byte(input)); err != nil {
		return "", err
	}

	// 计算哈希并获取其字节表示
	hashBytes := hasher.Sum(nil)

	// 将字节序列转换为十六进制字符串
	return fmt.Sprintf("%x", hashBytes), nil
}

func ValidateSignature(time int64, appID, secret, signature string) bool {
	// 生成签名
	generatedSignature, err := GenerateSHA1Signature(time, appID, secret)
	if err != nil {
		return false
	}

	// 比较生成的签名和传入的签名
	return signature == generatedSignature
}
