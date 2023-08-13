package src

import "testing"

func TestKeyGen(t *testing.T) {
	type args struct {
		base     uint64
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "0",
			args: args{
				20, "../data/dice_test.txt",
			},
			want: 6_615_175_401_215,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := KeyGen(tt.args.base, tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyGen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("KeyGen() got = %v, want %v", got, tt.want)
			}
		})
	}
}
