package main

import (
	"crypto/md5"
	"encoding/hex"
)

func CreatShortUrl(longUrl string) string {
	hash := md5.Sum([]byte(longUrl))
	shortUrl := hash[:5]
	return hex.EncodeToString(shortUrl)
}
