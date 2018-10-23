package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	First string //capital as we are exporting
	Age   int
}

//https://mholt.github.io/json-to-go/
type Person12 struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	u1 := User{
		First: "James",
		Age:   32,
	}

	u2 := User{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := User{
		First: "M",
		Age:   54,
	}

	Users := []User{u1, u2, u3}

	fmt.Println(Users)
	bs2, err := json.Marshal(Users)
	if err != nil {
		fmt.Println("there was error")
	} else {
		fmt.Println("no errors detected, json marshalled")
	}
	fmt.Println(bs2)
	fmt.Println(string(bs2))
	fmt.Println(string(bs2), "json marshaled")
	fmt.Println("---------------")
	s6 := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s6)
	bs6 := []byte(s6)
	var bsUsers []Person12 //bsUsers := []Person12{} ?
	err = json.Unmarshal(bs6, &bsUsers)
	if err != nil {
		fmt.Println("there were errors unmarshaling")
	} else {
		fmt.Println("no probs unmarshalling")
	}
	fmt.Println(bs6)
	fmt.Println(string(bs6))

	for i, Person12 := range bsUsers {
		fmt.Println("person #", i)
		fmt.Println("\t\t", Person12.First, "age: ", Person12.Age)
		for _, saying := range Person12.Sayings {
			fmt.Println("\t\t", saying)
		}
	}
	// your code goes here
}
