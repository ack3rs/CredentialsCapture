package credentials

import (
	"testing"
)

func TestFirstName(t *testing.T) {

	U := User{Firstname: "", Lastname: "Ackroyd", Country: "UK", Email: "mark@ackroyd.net"}
	err := U.Save()
	if err == nil {
		t.Error("User Validation TestFirstName Failed")
	}
}

func TestLastName(t *testing.T) {

	U := User{Firstname: "Mark", Lastname: "", Country: "UK", Email: "mark@ackroyd.net"}
	err := U.Save()
	if err == nil {
		t.Error("User Validation TestLastName Failed")
	}
}

func TestCountry(t *testing.T) {

	U := User{Firstname: "Mark", Lastname: "Ackroyd", Country: "", Email: "mark@ackroyd.net"}
	err := U.Save()
	if err == nil {
		t.Error("User Validation TestCountry Failed")
	}
}

func TestEmail(t *testing.T) {

	U := User{Firstname: "Mark", Lastname: "Ackroyd", Country: "UK", Email: ""}
	err := U.Save()
	if err == nil {
		t.Error("User Validation TestEmail Failed")
	}

	U = User{Firstname: "Mark", Lastname: "Ackroyd", Country: "UK", Email: "NOT VALID EMAIL"}
	err = U.Save()
	if err == nil {
		t.Error("User Validation TestEmail Failed")
	}

}
