package main

import "fmt"

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	for _, a := range s.name {
		intA := int(a)
		if 97 <= intA && intA <= 122 {
			nameEncode += string(rune(122 - (intA - 97)))
		} else if 65 <= intA && intA <= 90 {
			nameEncode += string(rune(90 - (intA - 65)))
		} else {
			nameEncode += string(a)
		}
	}
	return nameEncode
}

func (s *student) Decode() string {
	return s.Encode()
}

func main() {
	var (
		menu int
		a           = student{}
		c    Chiper = &a
	)

	fmt.Println("[1] Encrypt \n[2] Decrypt \nChoose Your Menu > ")
	fmt.Scan(&menu)

	switch menu {
	case 1:
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student's Name " + a.name + " is : " + c.Encode())
	case 2:
		fmt.Print("\nInput Student's Encode Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of Student's Name " + a.name + " is : " + c.Decode())
	default:
		fmt.Println("Wrong input name menu")
	}
}
