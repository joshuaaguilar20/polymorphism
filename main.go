package main

import "fmt"

//polymorphism
//Conversion aka casting but not called casting
//Assertion
//Interface says hey baby if you got this method than your my type :)
type player struct {
	name     string
	age      int8
	score    int64
	loggedIn bool
}

type ai struct {
	player
	setting string
	mode    string
}

type human interface {
	//Allows Value to Be More Than one Type
	changeScore()
	//Adding Change Score to Interface,
	//Adds Stucts as Value on human, n
}

func IcallHuman(h human) {
	//Can take AI or player as both are attached to the type human
	//Assertion
	switch h.(type) {
	case player:
		fmt.Println("I was Passed into Interface Human \n", h.(player).name)
	case ai:
		fmt.Println(" KILL ALL HUMANS\n", h.(ai).name)
	}
	
}

func (s player) changeScore() {
	fmt.Printf("%s", s.score)
}

func (a ai) changeScore() {
	fmt.Printf("%s", a.score)
}

//Methods
func (a ai) changeMode() string {
	if a.mode != "" {
		a.mode = "Hard"
		fmt.Printf("%s", a.mode)
		return a.mode
	}
	return `error`
}
func main() {
	p1 := player{
		name:     "Joshua",
		age:      29,
		score:    10201,
		loggedIn: true,
	}
	ai := ai{
		player: player{
			name:     "AI",
			age:      100,
			score:    1020121,
			loggedIn: true,
		},
		setting: "Hard",
		mode:    "Easy",
	}
	IcallHuman(ai)
	IcallHuman(p1)
	conversion()
}

type DefNotAstring int

func conversion() {
	var x DefNotAstring = 120
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int64
	y = int64(x)
	//Coverting Type HotDog to Type Int.
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
