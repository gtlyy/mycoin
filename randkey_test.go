package mycoin

import (
	"os"
	"sync"
	"testing"
)

// 测试 randomChars 函数
func TestRandomChars(t *testing.T) {
	const n = 10                // 生成的字符数
	words := "0123456789abcdef" // 字符集
	expectedLength := n         // 期望的输出长度

	output, err := RandomChars(n, words)
	if err != nil {
		t.Fatalf("Error generating random chars: %v", err)
	}

	// 检查输出长度
	if len(output) != expectedLength {
		t.Errorf("Expected output length %d, got %d", expectedLength, len(output))
	}

	// 检查输出字符是否在字符集内
	for _, char := range output {
		if !containsRune(words, char) {
			t.Errorf("Character %c not in expected words", char)
		}
	}
}

// 测试 randomBtcPrivateKey 函数
func TestRandomBtcPrivateKey(t *testing.T) {
	const length = 64           // BTC 私钥的标准长度
	words := "0123456789abcdef" // 字符集

	privateKey, err := RandomBtcPrivateKey(length)
	if err != nil {
		t.Fatalf("Error generating BTC private key: %v", err)
	}
	t.Log(privateKey)

	// 检查输出长度
	if len(privateKey) != length {
		t.Errorf("Expected private key length %d, got %d", length, len(privateKey))
	}

	// 检查输出字符是否在字符集内
	for _, char := range privateKey {
		if !containsRune(words, char) {
			t.Errorf("Character %c not in expected words", char)
		}
	}
}

// 测试 crack
func TestCrack(t *testing.T) {
	addrRight := "16HUcQitGbVcFtXCkV6W9VkSEYfJiaZW1f"
	const iterations = 10000 // 每个协程的循环次数
	const numWorkers = 10    // 协程数量
	var wg sync.WaitGroup

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				// 生成随机 BTC 私钥
				privateKey, _ := RandomBtcPrivateKey(64)

				// 根据私钥生成地址
				addr := ToAddress2To9(privateKey)

				// 检查地址是否匹配
				if addr == addrRight {
					// 找到匹配的私钥，写入文件
					err := WriteToFile("right.txt", privateKey)
					if err != nil {
						t.Errorf("Error writing to file: %v", err)
					}
					t.Logf("Found matching private key: %s", privateKey)

					// 立即退出程序
					os.Exit(0)
				}
			}
		}()
	}

	wg.Wait() // 等待所有协程完成
}
