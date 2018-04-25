package greeting

import (
	"fmt"
	)

const(
)

type Salutation struct {
	Name string
	Greeting string
}

type Renameable interface {
	Rename(newName string)
}

type Salutations []Salutation

type Printer func(string)()

func (salutations Salutations)  ChannelGreeter(c chan Salutation) {
	for _, s := range salutations {
		c <- s
	}
	close(c)
}

func (saluation *Salutation) Write(p []byte)(n int, err error){
	s := string(p)
	saluation.Rename(s)
	n = len(s)
	err = nil
	return
}

func (saluation *Salutation) Rename(newName string){
	saluation.Name = newName
}

func (saluations Salutations) Greet(do Printer, isFormal bool)  {

	for _, s := range saluations{

		_message, _altMsg := CreateMessage(s.Name, s.Greeting)
		if isFormal {
			do(_message)
		}else {
			do(_altMsg)
		}
	}
}

func GetPrefix(name string) (prefix string) {

	prefixMap := map[string]string{
		"Jo":     "Mrs.",
		"Danny":  "Dr",
		"Danhua": "Dr",
		"Aria":   "Miss",
	}

	prefixMap["Aria"] = "Jr"
	delete(prefixMap, "Danhua")

	if value, exists := prefixMap[name]; exists {
		return value
	}
	return "Dude"
	return prefixMap[name]
	/*
	switch {
	case name == "Jo": prefix = "Mrs."
	case name == "Danny" || name == "Danhua": prefix = "Dr."
	case name == "Aria": prefix = "Miss."
	case len(name) > 10 : prefix = "Lord"
	default: prefix = "Person"
	}
*/
}

func GetPrefixBasedOnType(x interface{}) (prefix string){
	switch x.(type){
	case int : prefix = "Mrs."
	case string : prefix = "Dr."
	case Salutation : prefix = "Miss."
	default: prefix = "Person"
	}
	return
}

func CreateMessage(name string, greeting ...string)  (message, alt string){
	prefix := GetPrefix(name);
	message = greeting[0] + " " + prefix + " " + name
	alt = "'s up dude " + name
	return
}

func Print(s string){
	fmt.Print(s)
}

func Println(s string){
		fmt.Println(s)
}

func PrintlnCustom(cust string) Printer{
	return func(s string){
			fmt.Println(cust + " " + s)
	}
}