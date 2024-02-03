package utilx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertType(t *testing.T) {
	var val1 any = Pointer("val1")
	// normal case
	pStr, err := ConvertType[*string](val1)
	assert.Nil(t, err)
	assert.Equal(t, "val1", *pStr)

	// error case
	zeroVal, err := ConvertType[string](val1)
	assert.NotNil(t, err)
	assert.Equal(t, "", zeroVal)
}
