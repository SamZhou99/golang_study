package img

import (
	"fmt"
	"strings"
)

// 全局替换器，程序启动时初始化一次，重复调用性能最优
var cardReplacer = strings.NewReplacer(
	"a", "♠",
	"b", "♥",
	"c", "♣",
	"d", "♦",
	"13", "K",
	"12", "Q",
	"11", "J",
	"10", "10",
	"1", "A",
)

// ConvertCardString 转换扑克牌字符串
// 替换规则：
//
//	花色：a→♠️  b→♥️  c→♣️  d→♦️
//	牌面：1→A  11→J  12→Q  13→K
func ConvertCardString(s string) string {
	return cardReplacer.Replace(s)
}

// 测试
func main() {
	// 测试示例
	fmt.Println(ConvertCardString("a1"))     // 输出：♠️A
	fmt.Println(ConvertCardString("b11"))    // 输出：♥️J
	fmt.Println(ConvertCardString("c12"))    // 输出：♣️Q
	fmt.Println(ConvertCardString("d13"))    // 输出：♦️K
	fmt.Println(ConvertCardString("b11 a1")) // 输出：♠️A ♥️J
}
