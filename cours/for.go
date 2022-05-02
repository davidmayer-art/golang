package cours

import "fmt"

func While() {
	i := 0

	for {
		if i > 10 {
			break
		}
		//algo
		i++
	}
}

func doWhile() {
	i := 0

	for {
		//Algo
		i++
		if i > 10 {
			break
		}
	}
}

func boucleFor() {
	for i := 0; i < 10; i++ {

	}
}

func boucleForEach() {
	a := []string{"toto", "Tata", "Titi"}

	for index := range a {
		fmt.Println(index)
	}
	for index, value := range a {
		fmt.Println(index, value)
	}
	for _, value := range a {
		fmt.Println(value)
	}
}
