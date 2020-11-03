package utils

import (
	testify "github.com/stretchr/testify/assert"
	"testing"
)

func Test_getPinyin(t *testing.T){
	pinyin := getPinyin("我爱中国")
	assert := testify.New(t)
	assert.Equalf("wo ai zhong guo", pinyin, "should be %s, but %s was given", pinyin, pinyin)
}