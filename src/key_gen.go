package src

import (
	"bufio"
	"io"
	"math/big"
	"os"
	"strconv"
)

func KeyGen(base uint64, fileName string) (*big.Int, int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, 0, err
	}

	bu := bufio.NewReaderSize(f, 1024)

	ans := big.NewInt(0)
	unit := big.NewInt(1)
	bigBase := big.NewInt(int64(base))

	lines := 0
	for {
		line, _, err := bu.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, 0, err
		}

		lines++
		num, err := strconv.Atoi(string(line))
		bigNum := big.NewInt(int64(num))
		//ans += unit * (uint64(num) % base)
		_, r := new(big.Int).DivMod(bigNum, bigBase, new(big.Int))
		ans.Add(ans, new(big.Int).Mul(unit, r))

		unit.Mul(unit, bigBase)
	}

	return ans, lines, nil
}
