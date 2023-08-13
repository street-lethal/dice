package src

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

func (s StringGenerator) Gen(key uint64, length int) string {
	var r int
	var q = key
	var str string

	size := uint64(len(s.chars))

	for i := 0; i < length; i++ {
		r = int(q % size)
		q /= size

		str += string(s.chars[r])
	}

	return str
}

func (s StringGenerator) NecessarySize(length, base int) (
	necessarySize int, fullSize uint64,
) {
	charLength := uint64(len(s.chars))
	fullSize = 1
	for i := 0; i < length; i++ {
		fullSize *= charLength
	}

	necessarySize = 0
	test := uint64(1)
	for test < fullSize {
		test *= uint64(base)
		necessarySize++
	}

	return
}
