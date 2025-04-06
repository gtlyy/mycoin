package mycoin

import (
	"crypto/sha256"
	"encoding/hex"

	// "fmt"
	"hash"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

// func main() {
// 	// 解码私钥
// 	privateKeyHex := "18e14a7b6a307f426a94f8114701e7c8e774e7f9a47e2c2035db29a206321725" // 这是一个示例私钥

// 	btcAddress_,:= ToAddress2To9(privateKeyHex)
// }

// Calculate the hash of hasher over buf.
// from btcutil hash160.go
func calcHash(buf []byte, hasher hash.Hash) []byte {
	hasher.Write(buf)
	return hasher.Sum(nil)
}

// 第二步，使用椭圆曲线加密算法（ECDSA-SECP256k1）计算私钥所对应的非压缩公钥（共65字节，1字节0x04，32字节为x坐标，32字节为y坐标）。
func ToPublicKeyHex(privateKeyHex string) string {
	// 解码私钥
	privateKeyBytes, _ := hex.DecodeString(privateKeyHex)
	privateKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privateKeyBytes)
	// 计算非压缩公钥
	publicKey := privateKey.PubKey()
	publicKeyHex := hex.EncodeToString(publicKey.SerializeUncompressed())
	return publicKeyHex
}

// 第三步，计算公钥的SHA-256哈希值
func ToSha256Hex(publicKeyHex string) string {
	pubKeyBytes, _ := hex.DecodeString(publicKeyHex)
	h := sha256.New()
	h.Write(pubKeyBytes)
	pubSha256Bytes := h.Sum(nil)
	pubSha256Hex := hex.EncodeToString(pubSha256Bytes)
	return pubSha256Hex
}

// 第四步，计算上一步哈希值的RIPEMD-160哈希值
func ToRipemd160Hex(pubSha256Hex string) string {
	pubSha256Bytes, _ := hex.DecodeString(pubSha256Hex)
	h := ripemd160.New()
	h.Write(pubSha256Bytes)
	pubRipemd160Bytes := h.Sum(nil)
	pubRipemd160Hex := hex.EncodeToString(pubRipemd160Bytes)
	return pubRipemd160Hex
}

// 第五步，在上一步结果之前加入地址版本号（如比特币主网版本号"0x00"）
func ToAddVersion(pubRipemd160Hex string) string {
	pubAddVersion := "00" + pubRipemd160Hex
	return pubAddVersion
}

// 第六步，计算上一步结果的SHA-256哈希值
func Tosha256FirstHex(pubAddVersion string) string {
	pubAddVersionBytes, _ := hex.DecodeString(pubAddVersion)
	sha256FirstBytes := calcHash(pubAddVersionBytes, sha256.New())
	sha256FirstHex := hex.EncodeToString(sha256FirstBytes)
	return sha256FirstHex
}

// 第七步，再次计算上一步结果的SHA-256哈希值
func Tosha256AgainHex(sha256FirstHex string) string {
	sha256FirstBytes, _ := hex.DecodeString(sha256FirstHex)
	sha256AgainBytes := calcHash(sha256FirstBytes, sha256.New())
	sha256AgainHex := hex.EncodeToString(sha256AgainBytes)
	return sha256AgainHex
}

// 第八步，取上一步结果的前4个字节（8位十六进制数）D61967F6，把这4个字节加在第五步结果的后面，作为校验（这就是比特币地址的16进制形态）
func ToaddressHex(pubAddVersion, sha256AgainHex string) string {
	addressHex := pubAddVersion + sha256AgainHex[:8]
	return addressHex
}

// 第九步，用base58表示法变换一下地址（这就是最常见的比特币地址形态）
func ToaddressBTC(addressHex string) string {
	addressBytes, _ := hex.DecodeString(addressHex)
	addressBTC := base58.Encode(addressBytes)
	return addressBTC
}

// 第二至九步，写在一起，如下：
func ToAddress2To9(privateKeyHex string) string {
	// 第二步
	// 解码私钥
	privateKeyBytes, _ := hex.DecodeString(privateKeyHex)
	privateKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privateKeyBytes)
	// 计算非压缩公钥
	publicKey := privateKey.PubKey()
	publicKeyHex := hex.EncodeToString(publicKey.SerializeUncompressed())
	// fmt.Println(publicKeyHex)

	// 第三步，计算公钥的SHA-256哈希值
	pubKeyBytes, _ := hex.DecodeString(publicKeyHex)
	h := sha256.New()
	h.Write(pubKeyBytes)
	pubSha256Bytes := h.Sum(nil)
	pubSha256Hex := hex.EncodeToString(pubSha256Bytes)
	// fmt.Println(pubSha256Hex)

	// 第四步，计算上一步哈希值的RIPEMD-160哈希值
	pubSha256Bytes2, _ := hex.DecodeString(pubSha256Hex)
	h2 := ripemd160.New()
	h2.Write(pubSha256Bytes2)
	pubRipemd160Bytes := h2.Sum(nil)
	pubRipemd160Hex := hex.EncodeToString(pubRipemd160Bytes)
	// fmt.Println(pubRipemd160Hex)

	// 第五步，在上一步结果之前加入地址版本号（如比特币主网版本号"0x00"）
	pubAddVersion := "00" + pubRipemd160Hex
	// fmt.Println(pubAddVersion)

	// 第六步，计算上一步结果的SHA-256哈希值
	pubAddVersionBytes, _ := hex.DecodeString(pubAddVersion)
	sha256FirstBytes := calcHash(pubAddVersionBytes, sha256.New())
	// sha256FirstHex := hex.EncodeToString(sha256FirstBytes)
	// fmt.Println(sha256FirstHex)

	// 第七步，再次计算上一步结果的SHA-256哈希值
	// sha256FirstBytes, _ = hex.DecodeString(sha256FirstHex)
	sha256AgainBytes := calcHash(sha256FirstBytes, sha256.New())
	sha256AgainHex := hex.EncodeToString(sha256AgainBytes)
	// fmt.Println(sha256AgainHex)

	// 第八步，取上一步结果的前4个字节（8位十六进制数）D61967F6，把这4个字节加在第五步结果的后面，作为校验（这就是比特币地址的16进制形态）
	addressHex := pubAddVersion + sha256AgainHex[:8]
	// fmt.Println(addressHex)

	// 第九步，用base58表示法变换一下地址（这就是最常见的比特币地址形态）
	addressBytes, _ := hex.DecodeString(addressHex)
	addressBTC := base58.Encode(addressBytes)
	// fmt.Println(addressBTC)
	return addressBTC
}
