package strx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBlank(t *testing.T) {
	blankStr := " "
	emptyStr := ""
	notBlankStr := " fef d "

	assert.True(t, IsBlank(blankStr))
	assert.True(t, IsBlank(emptyStr))
	assert.False(t, IsBlank(notBlankStr))
}

func TestIsEmpty(t *testing.T) {
	blankStr := " "
	emptyStr := ""
	notBlankStr := " fef d "

	assert.False(t, IsEmpty(blankStr))
	assert.True(t, IsEmpty(emptyStr))
	assert.False(t, IsEmpty(notBlankStr))
}
