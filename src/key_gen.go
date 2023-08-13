package src

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func KeyGen(base uint64, fileName string) (uint64, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}

	bu := bufio.NewReaderSize(f, 1024)

	var ans uint64 = 0
	var unit uint64 = 1

	for {
		line, _, err := bu.ReadLine()
		if err == io.EOF {
			break
		}

		num, err := strconv.Atoi(string(line))
		if err != nil {
			return 0, err
		}

		ans += unit * (uint64(num) % base)
		unit *= base
	}

	return ans, nil
}
