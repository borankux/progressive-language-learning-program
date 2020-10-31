package utils

import "github.com/mozillazg/go-pinyin"

func getPinyin(word string) [][]string {
	profile := pinyin.NewArgs()
	profile.Style = pinyin.Tone
	return pinyin.Pinyin(word, profile)
}

