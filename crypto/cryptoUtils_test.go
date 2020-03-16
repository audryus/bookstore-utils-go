package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMd5EmptyOrNil(t *testing.T) {
	md5 := GetMd5("")
	assert.NotNil(t, md5)
	assert.EqualValues(t, "", md5)

	var newString string

	md5 = GetMd5(newString)
	assert.NotNil(t, md5)
	assert.EqualValues(t, "", md5)
}
func TestGetMd5(t *testing.T) {
	md5 := GetMd5("my test password")
	
	assert.NotNil(t, md5)
	assert.EqualValues(t, "4ae6fe0e355169f70c7f6754ebf7b597", md5)

}