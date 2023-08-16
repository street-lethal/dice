package src

import "math/big"

type StringGenerator struct {
	chars []byte
}

func NewStringGenerator(charKinds string) StringGenerator {
	var chars []byte
	var tmp byte
	for _, charKind := range charKinds {
		duplicate := false
		for _, char := range chars {
			if char == byte(charKind) {
				duplicate = true
				break
			}
		}
		if duplicate {
			continue
		}

		strCharKind := string(charKind)
		if strCharKind == "A" {
			for tmp = 0x41; tmp <= 0x5a; tmp++ {
				chars = append(chars, tmp)
			}
		} else if strCharKind == "a" {
			for tmp = 0x61; tmp <= 0x7a; tmp++ {
				chars = append(chars, tmp)
			}
		} else if strCharKind == "0" {
			for tmp = 0x30; tmp <= 0x39; tmp++ {
				chars = append(chars, tmp)
			}
		} else if strCharKind == "#" ||
			strCharKind == "$" ||
			strCharKind == "%" ||
			strCharKind == "&" ||
			strCharKind == "_" ||
			strCharKind == "-" {
			chars = append(chars, byte(charKind))
		}
	}

	return StringGenerator{
		chars: chars,
	}
}

func (s *StringGenerator) ForbidCharacters(chars string) {
	for _, forbidden := range []byte(chars) {
		for i, char := range s.chars {
			if char == forbidden {
				s.chars = append(s.chars[:i], s.chars[i+1:]...)
			}
		}
	}
}

func (s StringGenerator) Gen(key *big.Int, length int) string {
	var r *big.Int
	var q = key
	var str string

	size := uint64(len(s.chars))

	for i := 0; i < length; i++ {
		q, r = new(big.Int).DivMod(
			q, big.NewInt(int64(size)), new(big.Int),
		)

		str += string(s.chars[r.Int64()])
	}

	return str
}

func (s StringGenerator) NecessarySize(length, base int) (
	necessarySize int, fullSize *big.Int,
) {
	charLength := big.NewInt(int64(len(s.chars)))
	fullSize = big.NewInt(1)
	for i := 0; i < length; i++ {
		fullSize.Mul(fullSize, charLength)
	}

	necessarySize = 0
	test := big.NewInt(1)
	bigBase := big.NewInt(int64(base))
	for test.Cmp(fullSize) < 0 {
		test.Mul(test, bigBase)
		necessarySize++
	}

	return
}
