package main

import (
	"crypto/md5"
	"encoding/hex"
)

func CreatShortUrl(longUrl string) string {
	hash := md5.Sum([]byte(longUrl))
	shortUrl := hash[:]
	return hex.EncodeToString(shortUrl)[:5]
}
