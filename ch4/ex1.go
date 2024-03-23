package ch4

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"ttharsh.gobook/ch2/popcount"
)

func ShaDiff(x1, x2 string) uint16 {
	if x1 == x2 {
		return 0
	}

	shax1 := sha256.Sum256([]byte(x1))
	shax2 := sha256.Sum256([]byte(x2))

	fmt.Printf("SHA256 of `%s` = %x\n", x1, shax1)
	fmt.Printf("SHA256 of `%s` = %x\n", x2, shax2)

	var t1, t2 [4]uint64

	for i := 0; i < 4; i++ {
		start := i * 8
		end := (i + 1) * 8

		t1[i] = binary.LittleEndian.Uint64(shax1[start:end])
		t2[i] = binary.LittleEndian.Uint64(shax2[start:end])
	}

	return bitDiff(&t1, &t2)
}

func bitDiff(b1, b2 *[4]uint64) uint16 {
	var sum uint16
	for i := 0; i < len(b1); i++ {
		xorb := b1[i] ^ b2[i]
		sum += uint16(popcount.PopCountv2(xorb))
	}
	return sum
}
