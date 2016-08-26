/*
生成一个ID字符串
前缀字符串(默认XY)+时间字符串（10个字符，共60bit）+若干个字符的字符串(默认3个字符)+校验位（1个字符）

注：每秒有10亿纳秒，再加3个字符的字符（共18bit），所能表达的空间超过250万亿/每秒，重复的概率可以忽略不计。
*/
package idGenerate

import (
	"sync"
	"time"
)

var (
	// ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-
	base64_table string = "wxyz01ABCDIJKLMNOlmn456EFGHPQRSZabcdefgTUVWXYhijopq_rstuv23k789-"

	// ID的前缀
	id_prefix string = "XY"

	// 当前的序列号
	// 初始值随机生成, 然后顺序加1
	curr_index int64

	mutex sync.Mutex
)

const (
	WORD_BITS uint8 = 6                   // 6个二进制位组成一个64进制字符
	WORD_MASK int64 = 63                  // 2^6 - 1
	BASE_NANO int64 = 1469631823539355011 // 2016-07-27的time.Now().UnixNano()

	// ID的随机字符串的字符数
	// 注：该值不能超过5
	ID_RAND_LEN uint8 = 3

	// ID随机字符串对应整数的最大值
	ID_RAND_MAX int64 = 1<<(3*6) - 1
)

// 初始化, 可以设置特殊的前缀字符串
// @param string prefix 最后生成的随机字符串的前缀
func Init(prefix string) {
	id_prefix = prefix
}

// 获取一个ID
func NextId() string {
	var (
		nano_time, rand_int  int64
		time_str, rand_str   string
		time_code, rand_code int64
		code                 byte
	)

	// 时间字符串
	nano_time = time.Now().UnixNano() - BASE_NANO
	time_str, time_code = int64_to_str(nano_time, 10)

	// 获取下一个随机值
	mutex.Lock()
	curr_index++
	if curr_index >= ID_RAND_MAX {
		curr_index = 0
	}
	rand_int = curr_index
	mutex.Unlock()

	// 随机字符串
	rand_str, rand_code = int64_to_str(rand_int, ID_RAND_LEN)

	// 计算校验码
	code = base64_table[(time_code+rand_code)&WORD_MASK]

	return id_prefix + time_str + rand_str + string(code)
}

// 将uint64转化为字符串
func int64_to_str(key int64, len uint8) (string, int64) {
	var (
		bytes [10]byte
		mask  int64
		code  int64 = 38 // 校验码有一个随机的初值
		i     uint8
	)

	for ; i < len; i++ {
		mask = key & WORD_MASK
		bytes[len-i-1] = base64_table[mask]
		key = key >> WORD_BITS
		code = (code + mask) & WORD_MASK
	}

	return string(bytes[:len]), code
}
