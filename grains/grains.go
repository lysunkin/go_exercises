package grains

import (
	"errors"
)

func Square(input int) (uint64, error) {
	if input < 1 || input > 64 {
		return 0, errors.New("invalid cell number")
	}

	var init uint64 = 1

	return init << (input-1), nil
}

func Total() uint64 {

	var result uint64 = 0 	

	for i:=0; i<64; i++ {
		var init uint64 = 1
		result += (init << i)
	}

	return result
}
