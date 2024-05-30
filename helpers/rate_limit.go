package helpers

import (
	"os"
	"strconv"
)

func GetRlValues() (int, int) {
	rlRate, rlBurst := os.Getenv("RL_RATE"), os.Getenv("RL_BURST")
	if (rlRate == "") || (rlBurst == "") {
		panic(ErrEnvConvert)
	}

	rlRateConv, err := strconv.Atoi(rlRate)
	if err != nil {
		panic(ErrEnvConvert)
	}

	rlBurstConv, err := strconv.Atoi(rlBurst)
	if err != nil {
		panic(ErrEnvConvert)
	}

	return rlRateConv, rlBurstConv
}
