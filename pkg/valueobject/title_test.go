package valueobject

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTitle(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name        string
		args        args
		wantAdTitle Title
		wantErr     bool
	}{
		{"empty title should error", args{""}, "", true},
		{"3 letters title should error", args{"err"}, "", true},
		{"6 or more letters should success", args{"Ad Title"}, "Ad Title", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAdTitle, err := NewTitle(tt.args.value)

			assert.Equal(t, tt.wantAdTitle, gotAdTitle)
			switch tt.wantErr {
			case true:
				assert.NotNil(t, err)
			case false:
				assert.Nil(t, err)
			}
		})
	}
}
