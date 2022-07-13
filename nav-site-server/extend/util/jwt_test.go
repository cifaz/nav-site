package util

import (
	"fmt"
	"testing"
)

func TestJWT_Make(t *testing.T) {
	jwt := JWT{}
	s, err := jwt.Make("admin", "nav-site-web", 7200)
	fmt.Println(s, err)
}

func TestJWT_Check(t *testing.T) {
	//app := conf.App

	jwt := JWT{}
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc3MTY2NTIsIm5hbWUiOiJhZG1pbiJ9.10FOjxfYCkLP5FvQ78BZcBQ2g5Ync9-vTGwgVQT1wqY"
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc3MTg2ODIsIm5hbWUiOiJhZG1pbiJ9.UbRmnvlpD8cUoZgUxzhudq3B3Aj7nWE9BAXYZITfips"
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc3MTg3NjgsIm5hbWUiOiJhZG1pbiJ9.1VoxIJY7GgtGqQzZf7fOEAiJB6hsNqz9dCSoUodHyXw"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc3MjcyMTksIm5hbWUiOiJhZG1pbiJ9.ujyGwZW0_pd3bd7fimfWFvKIUzPGtjvDkasx2OapKZE"
	check, err := jwt.Check(token, "nav-site-web")
	fmt.Println("check:", check, err)
}
