package src

import (
	"math/big"
	"testing"
)

func TestNewStringGenerator(t *testing.T) {
	type args struct {
		charKinds string
	}

	tests := []struct {
		name        string
		args        args
		wantedChars string
	}{
		{
			name:        "upper",
			args:        args{"A"},
			wantedChars: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		{
			name:        "lower",
			args:        args{"a"},
			wantedChars: "abcdefghijklmnopqrstuvwxyz",
		},
		{
			name:        "num",
			args:        args{"0"},
			wantedChars: "0123456789",
		},
		{
			name:        "upper + lower",
			args:        args{"Aa"},
			wantedChars: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
		},
		{
			name:        "upper + lower + num",
			args:        args{"Aa0"},
			wantedChars: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
		},
		{
			name:        "upper + lower + num + char",
			args:        args{"Aa0#$%&'()-+_"},
			wantedChars: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789#$%&-_",
		},
		{
			name:        "duplicates",
			args:        args{"Aa0Aa0"},
			wantedChars: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewStringGenerator(tt.args.charKinds)
			if string(got.chars) != tt.wantedChars {
				t.Errorf("got %v, wanted%v", string(got.chars), tt.wantedChars)
			}
		})
	}
}

func TestStringGenerator_Gen(t *testing.T) {
	largeKey, _ := new(big.Int).SetString("77767574737271706766656463626160", 8)
	tests := []struct {
		name  string
		chars []byte
		key   *big.Int
		size  int
		want  string
	}{
		{
			name:  "0",
			chars: []byte("012"),
			key:   big.NewInt(0),
			size:  3,
			want:  "000",
		},
		{
			name:  "1",
			chars: []byte("012"),
			key:   big.NewInt(1),
			size:  3,
			want:  "100",
		},
		{
			name:  "2",
			chars: []byte("012"),
			key:   big.NewInt(2),
			size:  3,
			want:  "200",
		},
		{
			name:  "3",
			chars: []byte("012"),
			key:   big.NewInt(3),
			size:  3,
			want:  "010",
		},
		{
			name:  "4",
			chars: []byte("012"),
			key:   big.NewInt(4),
			size:  3,
			want:  "110",
		},
		{
			name:  "8",
			chars: []byte("012"),
			key:   big.NewInt(8),
			size:  3,
			want:  "220",
		},
		{
			name:  "9",
			chars: []byte("012"),
			key:   big.NewInt(9),
			size:  3,
			want:  "001",
		},
		{
			name:  "10",
			chars: []byte("012"),
			key:   big.NewInt(10),
			size:  3,
			want:  "101",
		},
		{
			name:  "80",
			chars: []byte("012"),
			key:   big.NewInt(80),
			size:  3,
			want:  "222",
		},
		{
			name:  "81",
			chars: []byte("012"),
			key:   big.NewInt(81),
			size:  3,
			want:  "000",
		},
		{
			name:  "alphabet",
			chars: []byte("Ab3"),
			key:   big.NewInt(5),
			size:  3,
			want:  "3bA",
		},
		{
			name:  "large",
			chars: []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_-"),
			key:   largeKey,
			size:  20,
			want:  "mnopqrstuvwxyz_-0000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StringGenerator{
				chars: tt.chars,
			}
			if got := s.Gen(tt.key, tt.size); got != tt.want {
				t.Errorf("Gen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringGenerator_NecessarySize(t *testing.T) {
	large, _ := new(big.Int).SetString("3226266762397899821056", 10)
	tests := []struct {
		name              string
		chars             []byte
		length            int
		base              int
		wantNecessarySize int
		wantFullSize      big.Int
	}{
		{
			name:              "0",
			chars:             []byte("ABC"),
			length:            4,
			base:              8,
			wantNecessarySize: 3,
			wantFullSize:      *big.NewInt(81),
		},
		{
			name:              "0",
			chars:             []byte("ABC"),
			length:            4,
			base:              12,
			wantNecessarySize: 2,
			wantFullSize:      *big.NewInt(81),
		},
		{
			name:              "0",
			chars:             []byte("0123456789abcdef"),
			length:            8,
			base:              16,
			wantNecessarySize: 8,
			wantFullSize:      *big.NewInt(0x100_000_000),
		},
		{
			name:              "0",
			chars:             []byte("0123456789abcdef"),
			length:            8,
			base:              12,
			wantNecessarySize: 9,
			wantFullSize:      *big.NewInt(0x100_000_000),
		},
		{
			name:              "0",
			chars:             []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"),
			length:            12,
			base:              20,
			wantNecessarySize: 17,
			wantFullSize:      *large,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StringGenerator{
				chars: tt.chars,
			}
			gotNecessarySize, gotFullSize := s.NecessarySize(tt.length, tt.base)
			if gotNecessarySize != tt.wantNecessarySize {
				t.Errorf("NecessarySize() gotNecessarySize = %v, want %v", gotNecessarySize, tt.wantNecessarySize)
			}
			if gotFullSize.Cmp(&tt.wantFullSize) != 0 {
				t.Errorf("NecessarySize() gotFullSize = %v, want %v", gotFullSize, tt.wantFullSize)
			}
		})
	}
}
