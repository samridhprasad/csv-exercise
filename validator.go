package main

import (
	"regexp"
	"strconv"
)

type validationErr string

const (
	errHeaderMismatch validationErr = "record has an invalid number of fields: each record must have 5 fields to match with the CSV header"
	errInvalidID validationErr = "record has an invalid INTERNAL_ID: must be 8 digit positive int"
	errInvalidFname validationErr = "record has an invalid FIRST_NAME: expected 1-15 char max string, cannot be blank"
	errInvalidMname validationErr = "record has an invalid MIDDLE_NAME: expected 0-15 char max string, can be left blank"
	errInvalidLname validationErr = "record has an invalid LAST_NAME: expected 1-15 char max string, cannot be left blank"
	errInvalidPhone validationErr = "record has an invalid PHONE_NUM: expected format is ###-###-####"

)

func (err validationErr) Error() string { return string(err) }

// Reference: https://yourbasic.org/golang/regexp-cheat-sheet/
var (
	// 8 digit +ve integer, cannot be empty
	IdConstraints = regexp.MustCompile(`^[1-9]{8}$`)
	// 15 char max string, cannot be empty
	FirstNameConstraints = regexp.MustCompile(`^.{1,15}$`)
	// 15 char max string, can be empty
	MiddleNameConstraints = regexp.MustCompile(`^.{0,15}$`)
	// 15 char max string, cannot be empty
	LastNameConstraints = regexp.MustCompile(`^.{1,15}$`)
	// Format must align with ###-###-####
	PhoneConstraints = regexp.MustCompile(`^\d{3}-\d{3}-\d{4}$`)
)


func validateRecord(fields []string) (record Contact, err error) {
	if len(fields) != 5 {
		return record, errHeaderMismatch
	}
	strId := fields[0]
	id, _ := strconv.Atoi(fields[0])
	first := fields[1]
	middle := fields[2]
	last := fields[3]
	phone := fields[4]

	if !IdConstraints.MatchString(strId) {
		return record, errInvalidID
	} else if !FirstNameConstraints.MatchString(first) {
		return record, errInvalidFname
	} else if !MiddleNameConstraints.MatchString(middle) {
		return record, errInvalidMname
	} else if !LastNameConstraints.MatchString(last) {
		return record, errInvalidLname
	} else if !PhoneConstraints.MatchString(phone) {
		return record, errInvalidPhone
	} else {
		record.ID = id
		record.Name.First = first
		record.Name.Middle = middle
		record.Name.Last = last
		record.Phone = phone
	}
	return record, nil
}




