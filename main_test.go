package main

import (
	"regexp"
	"testing"
)

func Test_randChar(t *testing.T) {
	type args struct {
		length    int
		chars     []byte
		testCount int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"20 strings of 10 characters each",
			args{10, passwordChars, 20},
			10,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 1; i <= tt.args.testCount; i++ {
				chars, err := randChar(tt.args.length, tt.args.chars)
				if err != nil {
					t.Fatalf(err.Error())
				}
				if len(chars) != tt.want {
					t.Fatalf("generated string [%v] does not have [%v] characters", chars, tt.args.length)
				}
			}
		})
	}
}

func Test_newPassword(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"Password with 10 characters",
			args{10},
			10,
			false,
		},
		{
			"Password with 200 characters",
			args{200},
			200,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newPassword(tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("newPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			passwordLength := len(got)
			if passwordLength != tt.want {
				t.Errorf("newPassword() password [%v] has length [%v], want %v", got, passwordLength, tt.want)
			}
		})
	}
}

func Test_newSaltSHA512(t *testing.T) {
	type args struct {
		length    int
		testCount int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"10 salts",
			args{saltLength, 10},
			10,
			false,
		},
	}
	reString := `^\$6\$(\S+)$`
	re := regexp.MustCompile(reString)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 1; i <= tt.args.testCount; i++ {
				salt, err := newSaltSHA512(tt.args.length)
				if err != nil {
					t.Fatalf(err.Error())
				}
				if re.Match([]byte(salt)) {
					saltLengthWithOutPrefix := tt.args.length + 3
					if len(salt) != saltLengthWithOutPrefix {
						t.Fatalf("generated salt [%v] does not have [%v] characters", salt, saltLengthWithOutPrefix)
					}
				} else {
					t.Fatalf("generated salt [%v] does not match [%v]", salt, reString)
				}
			}
		})
	}
}
