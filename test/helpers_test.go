package test

import (
	"regexp"
	"testing"
)

func emailValidation(email string) bool {
	regex, _ := regexp.Compile(`^\w+([.]\w+)*@\w+([.]\w+)*\.\w+([.]\w+)*$`)
	return regex.MatchString(email)
}

func isBirthDateValid(s string) bool {
	regex, _ := regexp.Compile(`([12]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01]))`)
	return regex.MatchString(s)
}

func isPhoneValid(s string) bool {
	regex, _ := regexp.Compile(`^([0-9]{10,16})+$`)
	return regex.MatchString(s)
}

func TestEmailValidation(t *testing.T) {
	email1 := "xmortalx@noob.dev"
	// expected result is true
	if emailValidation(email1) {
		t.Log("expected result: email format is valid")
	} else {
		t.Error("unexpected result. test is failed")
	}


	// expected result is false
	email2 := "xmortalx@noob"
	if !emailValidation(email2) {
		t.Log("expected result: email format is not valid")
	} else {
		t.Error("unexpected result. test is failed")
	}

	email3 := "x_mortal_x@noob.dev"
	email4 := "x.mortal.x@noob.dev"
	email5 := "x-mortal-x@noob.dev"

	// email format check.
	t.Log(email1, ":", emailValidation(email1))
	t.Log(email2, ":", emailValidation(email2))
	t.Log(email3, ":", emailValidation(email3))
	t.Log(email4, ":", emailValidation(email4))
	t.Log(email5, ":", emailValidation(email5))
}

func TestBirthDateValidation(t *testing.T) {
	birthDate1 := "1999-12-31"
	// expected result is true
	if isBirthDateValid(birthDate1) {
		t.Log("expected result: birthdate format is valid")
	} else {
		t.Error("unexpected result. test is failed")
	}

	birthDate2 := "3999-12-31"
	// expected result is false
	if !isBirthDateValid(birthDate2) {
		t.Log("expected result: birthdate format is not valid")
	} else {
		t.Error("unexpected result. test is failed")
	}

	birthDate3 := "1999-13-31"
	birthDate4 := "199-12-31"
	birthDate5 := "1999-12-32"
	birthDate6 := "1999-1-01"
	birthDate7 := "1999-01-1"
	birthDate8 := "1999-01-01"

	// birthdate format check.
	t.Log(birthDate1, ":", isBirthDateValid(birthDate1))
	t.Log(birthDate2, ":", isBirthDateValid(birthDate2))
	t.Log(birthDate3, ":", isBirthDateValid(birthDate3))
	t.Log(birthDate4, ":", isBirthDateValid(birthDate4))
	t.Log(birthDate5, ":", isBirthDateValid(birthDate5))
	t.Log(birthDate6, ":", isBirthDateValid(birthDate6))
	t.Log(birthDate7, ":", isBirthDateValid(birthDate7))
	t.Log(birthDate8, ":", isBirthDateValid(birthDate8))

}

func TestPhoneValidation(t *testing.T) {
	phone1 := "081234567890"
	// expected result is true
	if isPhoneValid(phone1) {
		t.Log("expected result: phone format is valid")
	} else {
		t.Error("unexpected result. test is failed")
	}

	phone2 := "081234567890Z"
	// expected result is false
	if !isPhoneValid(phone2) {
		t.Log("expected result: phone format is not valid")
	} else {
		t.Error("unexpected result. test is failed")
	}

	phone3 := "081234"
	phone4 := "081234567891234567"

	// phone format check.
	t.Log(phone1, ":", isPhoneValid(phone1))
	t.Log(phone2, ":", isPhoneValid(phone2))
	t.Log(phone3, ":", isPhoneValid(phone3))
	t.Log(phone4, ":", isPhoneValid(phone4))


}