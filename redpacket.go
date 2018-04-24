// 观公众号文章有感 https://mp.weixin.qq.com/s/AIE33sdT2QI6UL8cs1kJCQ
package redpacket

import (
	"math/rand"
	"sort"
	"time"
)

// RedPacketV1 红包版本1. 根据  剩余红包金额/剩余可领取人数 * 2算出金额
func RedPacketV1(price int, numbers int) []int {
	restPrice, restNumbers := price, numbers

	rand.Seed(time.Now().UnixNano())
	packets := make([]int, numbers)
	for i := 0; i < numbers-1; i++ {
		value := rand.Intn(restPrice/restNumbers*2-1) + 1
		packets[i] = value
		restPrice -= value
		restNumbers--
	}

	packets[numbers-1] = restPrice

	return packets

}

// genateRandomNumbers 生成在区间(begin, end)内的不重复随机数
func genateRandomNumbers(begin, end, numbers int) []int {
	if end-begin < numbers {
		return []int{}
	}

	items := make([]int, end-begin)
	out := make([]int, end-begin)
	for i := 0; i < end-begin; i++ {
		items[i] = i + begin
	}

	for i := 0; i < numbers; i++ {
		restLength := len(items)
		index := rand.Intn(restLength)
		out[i] = items[index]
		items[index] = items[restLength-1]
		items = items[:restLength-1]
	}
	return out[:numbers]
}

// RedPacketV2 随机在一条线段上画numbers - 1个点, 切成numbers段
// 每段的长度即红包数
func RedPacketV2(price int, numbers int) []int {
	var currentPoint int
	packets := make([]int, numbers)

	randomNumbers := genateRandomNumbers(2, price, numbers-1)
	sort.Ints(randomNumbers)
	for i, number := range randomNumbers {
		packets[i] = number - currentPoint
		currentPoint = number
	}

	packets[numbers-1] = price - currentPoint
	return packets
}
