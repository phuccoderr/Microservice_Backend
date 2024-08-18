package basic

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAddOne(t *testing.T) {

	assert.Equal(t, AddOne(1), 2, "AddOne should be 3")
}
