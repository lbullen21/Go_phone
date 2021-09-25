package main

import (
	"regexp"
)

func normalize(phone string) string {
	//we want a number with no extra characters
	//this will creat new number string
	//this regex says we don't want to replace any of the digits 0-9
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone, "")
}
