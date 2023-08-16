package src

import (
	"math/big"
	"testing"
)

func TestKeyGen(t *testing.T) {
	largeKey, _ := new(big.Int).SetString("70810181890813363480520350", 10)
	type args struct {
		base     uint64
		fileName string
	}
	tests := []struct {
		name      string
		args      args
		want      *big.Int
		wantLines int
		wantErr   bool
	}{
		{
			name: "small",
			args: args{
				20, "../data/dice_test.txt",
			},
			want:      big.NewInt(6_615_175_401_215),
			wantLines: 10,
		},
		{
			name: "large",
			args: args{
				20, "../data/many_dice_test.txt",
			},
			want:      largeKey,
			wantLines: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotLines, err := KeyGen(tt.args.base, tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyGen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Cmp(tt.want) != 0 {
				t.Errorf("KeyGen() got = %v, want %v", got, tt.want)
			}
			if gotLines != tt.wantLines {
				t.Errorf("KeyGen() gotLines = %v, want %v", gotLines, tt.wantLines)
			}
		})
	}
}
