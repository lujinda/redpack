package redpacket

import (
	"fmt"
	"testing"
)

func sum(numbers []int) (total int) {
	for _, number := range numbers {
		total += number
	}
	return total
}

func testRedPacket(t *testing.T, vFunc func(int, int) []int) {
	tests := []struct {
		price   int
		numbers int
	}{
		{2000 * 100, 1},
		{1000 * 100, 2},
		{200 * 100, 30},
		{10 * 100, 4},
		{50 * 100, 9},
		{1 * 100, 20},
	}

	for _, test := range tests {
		packets := vFunc(test.price, test.numbers)
		if got := sum(packets); got != test.price {
			t.Errorf("%d块分%d分, 总额不等. got %d, want %d", test.price/100,
				test.numbers, got/100, test.price/100/100)
		} else {
			fmt.Println("红包:", packets)
		}
	}
}

func TestRedPacketV1(t *testing.T) {
	fmt.Println("V1")
	testRedPacket(t, RedPacketV1)
}

func TestRedPacketV2(t *testing.T) {
	fmt.Println("V2")
	testRedPacket(t, RedPacketV2)
}
