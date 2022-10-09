package helpers

import (
	"errors"
	"math"
	"os"
	"strings"
)

const (
	alphabet      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	largestuint64 = 18446744073709551615
)

func Base62Encode(number uint64) string {
	length := len(alphabet)
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(10)
	for ; number > 0; number = number / uint64(length) {
		encodedBuilder.WriteByte(alphabet[(number % uint64(length))])
	}

	return encodedBuilder.String()
}

func Base62Decode(encodedString string) (uint64, error) {
	var number uint64
	length := len(alphabet)

	for i, symbol := range encodedString {
		alphabeticPosition := strings.IndexRune(alphabet, symbol)
		if alphabeticPosition == -1 {
			return uint64(alphabeticPosition), errors.New("cannot find symbol in alphabet")
		}
		number += uint64(alphabeticPosition) * uint64(math.Pow(float64(length), float64(i)))
	}

	return number, nil
}

func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}

	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	if newURL == os.Getenv("DOMAIN") {
		return false
	}

	return true
}
