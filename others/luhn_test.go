package others

import (
	"errors"
	"testing"
)

func Test_luhn(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name string
		args args
		want bool
		wantErr error
	}{
		{
			name: "valid credit card number",
			args: args{
				number: "4539 3195 0343 6467",
			},
			want: true,
			wantErr: nil,
		},
		{
			name: "valid credit card number (no spaces)",
			args: args{
				number: "4539319503436467",
			},
			want: true,
			wantErr: nil,
		},
		{
			name: "invalid credit card number (num-numeric character)",
			args: args{
				number: "4539 3195 0343 646a",
			},
			want: false,
			wantErr: errors.New("all characters must be numbers"),
		},
		{
			name: "invalid credit card number",
			args: args{
				number: "4539 3195 0343 6468",
			},
			want: false,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := luhn(tt.args.number)
			if got != tt.want {
				t.Errorf("luhn() = %v, want %v", got, tt.want)
			}
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("luhn() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
