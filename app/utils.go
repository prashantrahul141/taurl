package app

import (
	"encoding/hex"
	"hash"
	"hash/fnv"
)

func GetFNVHash(text string) string {
	var hasher hash.Hash = fnv.New32a()
	hash := hasher.Sum([]byte(text))
	return hex.EncodeToString(hash[:4])
}

func shorten_url(original_url string) string {
	return GetFNVHash(original_url)
}
