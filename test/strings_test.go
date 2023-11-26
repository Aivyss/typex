package test

import (
	"github.com/aivyss/typex/strings"
	"github.com/aivyss/typex/testTool"
	"testing"
)

func TestIsBlank(t *testing.T) {
	blankStr := " "
	emptyStr := ""
	notBlankStr := " fef d "

	testTool.True(t, strings.IsBlank(blankStr))
	testTool.True(t, strings.IsBlank(emptyStr))
	testTool.False(t, strings.IsBlank(notBlankStr))
}

func TestIsEmpty(t *testing.T) {
	blankStr := " "
	emptyStr := ""
	notBlankStr := " fef d "

	testTool.False(t, strings.IsEmpty(blankStr))
	testTool.True(t, strings.IsEmpty(emptyStr))
	testTool.False(t, strings.IsEmpty(notBlankStr))
}
