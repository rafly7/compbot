package utils

import "log"

func Recover(err error) {
	if r := recover(); r != nil {
		log.Print(err)
	}
}
