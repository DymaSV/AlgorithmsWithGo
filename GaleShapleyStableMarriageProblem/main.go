package main

import (
	"fmt"
)

// Person is a struct that preset a man and a woman
type Person struct {
	id               int
	free             bool
	listOfPreference []int
}

//Pair is a struct that contains a man and a woman
type Pair struct {
	id    int
	man   *Person
	woman *Person
}

var pairs = []Pair{
	{id: 0, man: nil, woman: nil},
	{id: 1, man: nil, woman: nil},
	{id: 2, man: nil, woman: nil},
	{id: 3, man: nil, woman: nil},
}

func main() {
	men := []Person{
		Person{id: 0, free: true, listOfPreference: []int{0, 3, 1, 2}},
		Person{id: 1, free: true, listOfPreference: []int{2, 1, 0, 3}},
		Person{id: 2, free: true, listOfPreference: []int{2, 3, 1, 1}},
		Person{id: 3, free: true, listOfPreference: []int{2, 3, 0, 0}},
	}

	women := []Person{
		Person{id: 0, free: true, listOfPreference: []int{0, 2, 3, 1}},
		Person{id: 1, free: true, listOfPreference: []int{3, 2, 0, 1}},
		Person{id: 2, free: true, listOfPreference: []int{2, 3, 0, 1}},
		Person{id: 3, free: true, listOfPreference: []int{2, 3, 0, 1}},
	}

	pairs[0].woman = &women[0]
	pairs[1].woman = &women[1]
	pairs[2].woman = &women[2]
	pairs[3].woman = &women[3]

	// main program
	creatPairs(&men, &women)
	for _, i := range pairs {
		fmt.Printf("man - %v, woman - %v\n", i.man.id, i.woman.id)
	}
}

func creatPairs(men *[]Person, women *[]Person) {
	m, b := getFreeMan(men)
	if b {
		manPreferWomen := m.listOfPreference
		for _, idWoman := range manPreferWomen {
			w := &(*women)[idWoman]
			if w.free {
				pairs[idWoman].man = m
				w.free = false
				(*men)[m.id].free = false
				break
			} else {
				if proposeToWoman(w, m, &pairs) {
					(*men)[pairs[w.id].man.id].free = true
					pairs[w.id].man = m
					(*men)[m.id].free = false
					w.free = false
					break
				} else {
					continue
				}
			}
		}
	}
	_, ok := getFreeMan(men)
	if ok {
		creatPairs(men, women)
	}
}

func proposeToWoman(woman *Person, newMan *Person, pairs *[]Pair) bool {
	currentMan := (*pairs)[woman.id].man
	for _, wPrefer := range woman.listOfPreference {
		if wPrefer == newMan.id {
			return true
		}
		if wPrefer == currentMan.id {
			return false
		}
	}
	return false
}

func getFreeMan(men *[]Person) (*Person, bool) {
	for _, m := range *men {
		if m.free == true {
			return &m, true
		}
	}
	return nil, false
}
