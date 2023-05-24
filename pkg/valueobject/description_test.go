package valueobject

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDescription(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name              string
		args              args
		wantAdDescription Description
		wantErr           bool
	}{
		{"empty description should error", args{"Short description error"}, "", true},
		{"3 words description should error", args{""}, "", true},
		{"4 words description should success", args{"4 words description should success"}, "4 words description should success", false},
		{"right format description should success", args{"right format description should success"}, "right format description should success", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAdDescription, err := NewDescription(tt.args.value)

			assert.Equal(t, tt.wantAdDescription, gotAdDescription)
			switch tt.wantErr {
			case true:
				assert.NotNil(t, err)
			case false:
				assert.Nil(t, err)
			}
		})
	}
}
