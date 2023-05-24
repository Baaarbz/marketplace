package valueobject

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewId(t *testing.T) {
	const givenId = "574cc928-f4bd-11ed-ad0e-8a6a68a798d6"
	gotId, err := NewId(givenId)

	assert.Nil(t, err)
	assert.Equal(t, givenId, gotId.String())
}
func TestNewIdWrongFormat(t *testing.T) {
	id, err := NewId("wrong-id-format")

	assert.Error(t, err, "id: bad format")
	assert.Empty(t, id)
}
