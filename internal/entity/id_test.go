package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIDTypeValue(t *testing.T) {
	assert.Equal(t, "type", ID("type:value").Type())
	assert.Equal(t, "value", ID("type:value").ID())
}
