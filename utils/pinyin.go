package utils

import (
	"github.com/mozillazg/go-pinyin"
)

func getPinyin(word string) []string {
	profile := pinyin.NewArgs()
	profile.Style = pinyin.Tone
	py := pinyin.Pinyin(word, profile)
	var res []string
	for i:= range py {
		res = append(res, py[i][0])
	}
	return res
}

