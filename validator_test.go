package main

import (
	"reflect"
	"testing"
)

func TestValidateRecord(t *testing.T) {
	// Expected values for mock Contacts for tests
	var validRecord Contact
	validRecord.ID = 12345678
	validRecord.Name.First = "Samuel"
	validRecord.Name.Middle = "L"
	validRecord.Name.Last = "Jackson"
	validRecord.Phone = "179-066-7987"

	var validRecordWithoutMiddle Contact
	validRecordWithoutMiddle.ID = 91837121
	validRecordWithoutMiddle.Name.First = "Jonas"
	validRecordWithoutMiddle.Name.Last = "Kahnwald"
	validRecordWithoutMiddle.Phone = "929-280-1932"

	var invalidRecordWithBadID Contact
	invalidRecordWithBadID.ID = 0

	var invalidRecordWithBadFirstName Contact
	invalidRecordWithBadFirstName.Name.First = ""

	var invalidRecordWithBadMiddleName Contact
	invalidRecordWithBadFirstName.Name.Middle = ""

	var invalidRecordWithBadLastName Contact
	invalidRecordWithBadFirstName.Name.Last = ""

	var invalidRecordWithBadPhoneNum Contact
	invalidRecordWithBadPhoneNum.Phone = ""

	type csvline struct {
		fields []string
	}

	tests := []struct {
		testcase string
		csvline  csvline
		want     Contact
		catch    error
	}{
		{testcase: "Test for valid input for all fields", csvline: csvline{[]string{"12345678", "Samuel", "L", "Jackson", "179-066-7987"}}, want: validRecord},
		{testcase: "Test for invalid input with 16 digit ID", csvline: csvline{[]string{"9183712191837121", "Martha", "", "Nielson", "929-280-1932"}}, want: invalidRecordWithBadID, catch: errInvalidID},
		{testcase: "Test for invalid input with 0s in ID", csvline: csvline{[]string{"10101010", "Ulrich", "", "Nielson", "929-281-1932"}}, want: invalidRecordWithBadID, catch: errInvalidID},
		{testcase: "Test for invalid input with no ID", csvline: csvline{[]string{"", "Ulrich", "", "Nielson", "929-281-1932"}}, want: invalidRecordWithBadID, catch: errInvalidID},
		{testcase: "Test for invalid input with len(first name) exceeds 15", csvline: csvline{[]string{"91837115", "Schwarzkopfschmerzen", "Helge", "Nielson", "929-281-1992"}}, want: invalidRecordWithBadFirstName, catch: errInvalidFname},
		{testcase: "Test for invalid input without first name", csvline: csvline{[]string{"91837116", "", "", "Nielson", "929-281-1992"}}, want: invalidRecordWithBadFirstName, catch: errInvalidFname},
		{testcase: "Test for invalid input with len(middle name) exceeds 15", csvline: csvline{[]string{"91837115", "Helge", "Schwarzkopfschmerzen", "Nielson", "929-281-1992"}}, want: invalidRecordWithBadMiddleName, catch: errInvalidMname},
		{testcase: "Test for valid input without middle name", csvline: csvline{[]string{"91837121", "Jonas", "", "Kahnwald", "929-280-1932"}}, want: validRecordWithoutMiddle},
		{testcase: "Test for invalid input with len(last name) exceeds 15", csvline: csvline{[]string{"91837115", "John", "Helge", "Schwarzkopfschmerzen", "929-281-1992"}}, want: invalidRecordWithBadLastName, catch: errInvalidLname},
		{testcase: "Test for invalid input without last name", csvline: csvline{[]string{"91837116", "John", "Jacob", "", "929-281-1992"}}, want: invalidRecordWithBadLastName, catch: errInvalidLname},
		{testcase: "Test for invalid input with phone number pattern mismatch", csvline: csvline{[]string{"91837116", "John", "Jacob", "Wolf", "929-281-19-91"}}, want: invalidRecordWithBadPhoneNum, catch: errInvalidPhone},
		{testcase: "Test for invalid input without phone number", csvline: csvline{[]string{"91837116", "John", "Jacob", "Wolf", ""}}, want: invalidRecordWithBadPhoneNum, catch: errInvalidPhone},
		{testcase: "Test for invalid input with header mismatch", csvline: csvline{[]string{"Ulrich", "Nielson", "929-281-1932"}}, catch: errHeaderMismatch},
		{testcase: "Test for invalid input with empty fields", csvline: csvline{[]string{}}, catch: errHeaderMismatch},
	}

	for _, current := range tests {
		t.Run(current.testcase, func(t *testing.T) {
			resultVal, resultErr := validateRecord(current.csvline.fields)
			if !reflect.DeepEqual(resultVal, current.want) {
				t.Errorf("test failed: validateRecord() encountered = %v, expected = %v", resultVal, current.want)
			}
			if !reflect.DeepEqual(resultErr, current.catch) {
				t.Errorf("test failed: validateRecord() encountered = %v, expected = %v", resultErr, current.catch)

			}
		})
	}
}
