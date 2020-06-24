package main

import (
	"fmt"
	"sort"
)

// Starting with this code (https://play.golang.org/p/BVRZTdlUZ_), sort the []user by age, last
// Also sort each []string “Sayings” for each user
// print everything in a way that is pleasant

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}
type SortByAge []user
type SortByLast []user

func (a SortByAge) Len() int           { return len(a) }
func (a SortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (a SortByLast) Len() int           { return len(a) }
func (a SortByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	for _, v := range users {
		// Sort Saying
		sort.Strings(v.Sayings)
	}

	// your code goes here
	sort.Sort(SortByAge(users))
	sort.Sort(SortByLast(users))
	fmt.Printf("%+v", users)
}
