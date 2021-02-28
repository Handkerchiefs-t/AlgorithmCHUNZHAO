package Week_05

func reverseBits(num uint32) uint32 {
	res := uint32(0)

	for i := 0; i < 32; i++ {
		if (num >> i) & 1 == 1 {
			res = 1 << (31-i) | res
		}
	}

	return res
}
