package main

import (
	"danny.asb.co.nz/greeting"
	"fmt"
)

const(
	a = 1 <<(3*iota)
	b
	c
	d
	e
	f
	g
	h
	i
	j
)

func main() {
	x := 1

	for {
		x += x
		fmt.Println(x)

		if(x >= 102400){break;}
	}
}

func RenameToLion(r greeting.Renameable)  {
	r.Rename("Lion")
}
func main1() {
	//_slt := greeting.Solutation{"Danny","Greetings milord"}

	salutations := greeting.Salutations{
		{"Aria", "Hello"},
		{"Danny", "Hi"},
		{"Jo", "Allo"},
		{"Danhua", "Greetings"},
	}

	//slice = append(slice[:1], slice[2:]...)
	salutations[1].Rename("Daniel")
	RenameToLion(&salutations[2])

	c := make(chan greeting.Salutation)
	c2 := make(chan greeting.Salutation)

	go salutations.ChannelGreeter(c)
	go salutations.ChannelGreeter(c2)

	for {
		select {
		case s, ok := <-c:
			if ok {
				fmt.Println(s, ":1")
			} else {
				return
			}

		case s, ok := <-c2:
			if ok {
				fmt.Println(s, ":2")
			} else {
				return
			}

		default:
			fmt.Println("Waiting..")
		}
	}
}