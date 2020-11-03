package utils

import (
	testify "github.com/stretchr/testify/assert"
	"testing"
)

func Test_getPinyin(t *testing.T){
	pinyin := getPinyin("我爱中国")
	assert := testify.New(t)
	expected := []string{"wǒ","ài","zhōng","guó"}
	assert.Equalf(expected, pinyin, "should be %s, but %s was given", expected, pinyin)
}