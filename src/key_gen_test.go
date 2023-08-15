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
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "small",
			args: args{
				20, "../data/dice_test.txt",
			},
			want: big.NewInt(6_615_175_401_215),
		},
		{
			name: "large",
			args: args{
				20, "../data/many_dice_test.txt",
			},
			want: largeKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := KeyGen(tt.args.base, tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyGen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Cmp(tt.want) != 0 {
				t.Errorf("KeyGen() got = %v, want %v", got, tt.want)
			}
		})
	}
}
