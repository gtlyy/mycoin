package mycoin

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	_privateKeyHex   = "18e14a7b6a307f426a94f8114701e7c8e774e7f9a47e2c2035db29a206321725"
	_publicKeyHex    = "0450863AD64A87AE8A2FE83C1AF1A8403CB53F53E486D8511DAD8A04887E5B23522CD470243453A299FA9E77237716103ABC11A1DF38855ED6F2EE187E9C582BA6"
	_pubSha256Hex    = "600FFE422B4E00731A59557A5CCA46CC183944191006324A447BDB2D98D4B408"
	_pubRipemd160Hex = "010966776006953D5567439E5E39F86A0D273BEE"
	_pubAddVersion   = "00010966776006953D5567439E5E39F86A0D273BEE"
	_sha256FirstHex  = "445C7A8007A93D8733188288BB320A8FE2DEBD2AE1B47F0F50BC10BAE845C094"
	_sha256AgainHex  = "D61967F63C7DD183914A4AE452C9F6AD5D462CE3D277798075B107615C1A8A30"
	_addressHex      = "00010966776006953D5567439E5E39F86A0D273BEED61967F6"
	_addressBTC      = "16UwLL9Risc3QfPqBUvKofHmBQ7wMtjvM"
)

// 测试第二步，使用椭圆曲线加密算法（ECDSA-SECP256k1）计算私钥所对应的非压缩公钥（共65字节，1字节0x04，32字节为x坐标，32字节为y坐标）。
func Test_ToPublicKeyHex(t *testing.T) {
	assert.True(t, strings.ToUpper(ToPublicKeyHex(_privateKeyHex)) == _publicKeyHex)
}

// 测试第三步，计算公钥的SHA-256哈希值
func Test_ToSha256Hex(t *testing.T) {
	assert.True(t, strings.ToUpper(ToSha256Hex(_publicKeyHex)) == _pubSha256Hex)
}

// 第四步，计算上一步哈希值的RIPEMD-160哈希值
func Test_pubRipemd160Hex(t *testing.T) {
	s := strings.ToUpper(ToRipemd160Hex(_pubSha256Hex))
	assert.True(t, s == _pubRipemd160Hex)
}

// 测试第五步，在上一步结果之前加入地址版本号（如比特币主网版本号"0x00"）
func Test_pubAddVersion(t *testing.T) {
	assert.True(t, strings.ToUpper(ToAddVersion(_pubRipemd160Hex)) == _pubAddVersion)
}

// 测试第六步，计算上一步结果的SHA-256哈希值
func Test_sha256FirstHex(t *testing.T) {
	assert.True(t, strings.ToUpper(Tosha256FirstHex(_pubAddVersion)) == _sha256FirstHex)
}

// 测试第七步，再次计算上一步结果的SHA-256哈希值
func Test_sha256AgainHex(t *testing.T) {
	assert.True(t, strings.ToUpper(Tosha256AgainHex(_sha256FirstHex)) == _sha256AgainHex)
}

// 测试第八步，取上一步结果的前4个字节（8位十六进制数）D61967F6，把这4个字节加在第五步结果的后面，作为校验（这就是比特币地址的16进制形态）
func Test_addressHex(t *testing.T) {
	ah := ToaddressHex(_pubAddVersion, _sha256AgainHex)
	assert.True(t, strings.ToUpper(ah) == _addressHex)
}

// 测试第九步，用base58表示法变换一下地址（这就是最常见的比特币地址形态）
func Test_addressBTC(t *testing.T) {
	s := ToaddressBTC(_addressHex)
	assert.True(t, s == _addressBTC)
}

// 测试ToAddress2To9
func Test_ToAddress2To9(t *testing.T) {
	s := ToAddress2To9(_privateKeyHex)
	assert.True(t, s == _addressBTC)
}
