package puppy

import (
	"fmt"

	"github.com/Ng1n3/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! woof! woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp((Barks()))
}

func From1() {
  fmt.Println("I'm from version 1.0.0")
}