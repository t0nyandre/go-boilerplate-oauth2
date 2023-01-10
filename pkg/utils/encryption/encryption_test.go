package encryption

import (
	"reflect"
	"testing"
)

func Test_isHex(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isHex",
			args: args{s: []byte("c137e2f71944588efc0046c2faba9b7ed274d238da39c92509ab3038ed7709d5b1f356a8a37e62afd8a1da61dcf219c22ef5a84a81d8e1182ccf0cc619a20613")},
			want: true,
		},
		{
			name: "isNotHex",
			args: args{s: []byte("^+2ZN6NHu3mo0s2f9w0+-rEq^zwJRSFc8@Zc)G0jN!zCIE^nq@1q$tFMoQM7brnt")},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHex(tt.args.s); got != tt.want {
				t.Errorf("isHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSessionSecret(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		{
			name:    "getSessionSecret",
			want:    []byte{193, 55, 226, 247, 25, 68, 88, 142, 252, 0, 70, 194, 250, 186, 155, 126, 210, 116, 210, 56, 218, 57, 201, 37, 9, 171, 48, 56, 237, 119, 9, 213, 177, 243, 86, 168, 163, 126, 98, 175, 216, 161, 218, 97, 220, 242, 25, 194, 46, 245, 168, 74, 129, 216, 225, 24, 44, 207, 12, 198, 25, 162, 6, 19},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("SESSION_SECRET", "c137e2f71944588efc0046c2faba9b7ed274d238da39c92509ab3038ed7709d5b1f356a8a37e62afd8a1da61dcf219c22ef5a84a81d8e1182ccf0cc619a20613")
			got, err := getSessionSecret()
			if (err != nil) != tt.wantErr {
				t.Errorf("getSessionSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSessionSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}
