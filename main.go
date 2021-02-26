package main

import (
	"fmt"
	"reflect"
)

type employee struct {
	name      string
	age       int8
	salary    int
	rfc       string `json:"rfc"`
	gender    byte
	inability bool
}

func main() {
	var emp = employee{name: "Marco Díaz", age: 37, salary: 600, rfc: "VECJ880326", gender: 'M', inability: false}

	//Type, Kind and ValueOf
	fmt.Println("-- Type, Kind and ValueOf --")
	var rType = reflect.TypeOf(emp)   //type of interface
	var rKind = rType.Kind()          //kind of class
	var rValue = reflect.ValueOf(emp) //value
	fmt.Println(rType)
	fmt.Println(rKind)
	fmt.Println(rValue)

	var ptrE = &emp
	var ptrType = reflect.TypeOf(ptrE)
	fmt.Printf("Type: %s Kind: %s\n", ptrType.Elem(), ptrType.Kind()) //Kind: ptr (pointer)

	//Type of field
	fmt.Println("-- Type of field --")
	var eType = reflect.TypeOf(emp)
	for i := 0; i < eType.NumField(); i++ { // i < number of fields
		var field = eType.Field(i)
		fmt.Printf("Field Type: %s, Kind: %s\n", field.Name, field.Type.Kind()) //Name of field and kind
	}

	//Get tag
	fmt.Println("-- Get tag --")
	for i := 0; i < eType.NumField(); i++ {
		var field = eType.Field(i)
		if field.Tag != "" {
			fmt.Printf("Tag json: %s\n", field.Tag.Get("json"))
		}
	}

	//Change values
	fmt.Println("-- Change value --")

	var rEmployee = reflect.ValueOf(&emp) //Add pointer
	var newEmp = employee{name: "Ana Gomez", age: 41, salary: 700, rfc: "FRWJ950371", gender: 'F', inability: true}
	var rNewEmp = reflect.ValueOf(newEmp)
	rEmployee.Elem().Set(rNewEmp) //A set is applied to the rEmployee, so, the employee changes

	fmt.Println(emp)
	fmt.Printf("I'm %s, I´m %d years old and my RFC is: %s\n", emp.name, emp.age, emp.rfc)

}
