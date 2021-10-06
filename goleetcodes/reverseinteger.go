package main

import (
	"math"
)

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x *= sign
	}
	var d []int
	for x > 0 {
		d = append(d, x%10)
		x /= 10
	}
	if len(d) == 10 {
		sigDig := 0
		intMaxSigDigMag := 1_000_000_000
		for sigDig < 10 {
			intMaxSigDig := getIntMaxSigDig(intMaxSigDigMag, sign)
			//fmt.Printf("sigDig: %v, sigDigMag: %v. d[sigDig]: %v, MaxInt[sigDig]: %v\n", sigDig, intMaxSigDigMag, d[sigDig], intMaxSigDig)
			if d[sigDig] > intMaxSigDig {
				return 0
			} else if d[sigDig] == intMaxSigDig {
				sigDig += 1
				intMaxSigDigMag /= 10
			} else {
				break
			}
		}
	}
	reversed := 0
	i := 0
	for i < len(d) {
		reversed *= 10
		reversed += d[i]
		i++
	}
	return reversed * sign
}

func getIntMaxSigDig(intMaxSigDigMag int, sign int) int {
	if intMaxSigDigMag == 1 && sign < 0 {
		return (math.MaxInt32/intMaxSigDigMag)%10 - 1
	} else {
		return (math.MaxInt32 / intMaxSigDigMag) % 10
	}
}
