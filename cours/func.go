package cours

import "fmt"

func world(name string) (string, error) {

	if name == "" {
		return name, fmt.Errorf("Aucun name passé en argument")
	}
	return name, nil
}

func Hello(name string) (string, error) {

	if name == "" {
		name = "world"
	}

	return world(name)
}
