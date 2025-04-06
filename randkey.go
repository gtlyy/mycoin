package mycoin

import (
	"crypto/rand" // 随机性更好
	// "errors"
	"math/big"
	myrand "math/rand" // 速度更快，随机性较差
	"time"
)

// 随机生成字符并返回字符串
func RandomCharsMath(n int, words string) (string, error) {
	// if n <= 0 || len(words) == 0 {
	// 	return "", nil
	// }

	out := make([]byte, n)
	// 为每个调用创建独立的随机数生成器
	rng := myrand.New(myrand.NewSource(time.Now().UnixNano() + int64(n)))

	for i := 0; i < n; i++ {
		index := rng.Intn(len(words))
		out[i] = words[index]
	}
	return string(out), nil // 将字节数组转换为字符串并返回
}

// 随机生成字符并返回字符串
func RandomChars(n int, words string) (string, error) {
	// if n <= 0 {
	// 	return "", errors.New("n must be positive")
	// }

	length := len(words)
	// if length == 0 {
	// 	return "", errors.New("words cannot be empty")
	// }

	out := make([]byte, n)
	for i := 0; i < n; i++ {
		// 生成一个范围在 [0, length) 的随机数
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(length)))
		// if err != nil {
		// 	return "", err
		// }
		out[i] = words[index.Int64()]
	}
	return string(out), nil // 将字节数组转换为字符串并返回
}

// 随机生成BTC私钥 RandomChars s
func RandomBtcPrivateKey(length int) (string, error) {
	words := "0123456789abcdef"
	return RandomChars(length, words)
}

// 随机生成BTC私钥  RandomCharsMath
func RandomBtcPrivateKeyMath(length int) (string, error) {
	words := "0123456789abcdef"
	return RandomCharsMath(length, words)
}
