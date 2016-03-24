package popcount

// pc[i]はiのポピュレーションカウントです
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// 異なるビット数を数えます
func BitCount(x byte, y byte) int {
	var count = 0
	var i uint
	for i = 0; i < 8; i++ {
		if (int(pc[byte(x>>(i*8))])) != (int(pc[byte(y>>(i*8))])) {
			count++
		}
	}
	return count
}
