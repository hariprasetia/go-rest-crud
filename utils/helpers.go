package utils

import (
	"regexp"
	"time"
)

func Timestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func IsAlphabet(s string) bool {
	regex, _ := regexp.Compile(`(^[a-zA-Z]+$)`)
	return regex.MatchString(s)
}

func IsPhoneValid(s string) bool {
	regex, _ := regexp.Compile(`^([0-9]{7,16})+$`)
	return regex.MatchString(s)
}

func IsBirthDateValid(s string) bool {
	regex, _ := regexp.Compile(`([12]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01]))`)
	return regex.MatchString(s)
}

func IsEmailValid(s string) bool {
	regex, _ := regexp.Compile(`^\w+([.]\w+)*@\w+([.]\w+)*\.\w+([.]\w+)*$`)
	return regex.MatchString(s)
}