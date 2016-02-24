package popcount

// pc[i]はiのポピュレーションカウントです
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountはxのポピュレーションカウント（1が設定されているビット数）を返します。
func PopCount(x uint) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountRoot(x uint) int {
	var count = 0
	var i uint
	for i = 0; i < 8; i++ {
		count = count + int(pc[byte(x>>(i*8))])
	}
	return count
}

func PopCountBitShift(x uint) int {
	var count, i uint = 0, 0
	for i = 0; i < 64; i++ {
		count = count + (x >> i & 1)
	}
	return int(count)
}
