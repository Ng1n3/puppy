package puppy

import "github.com/GoesToEleven/dog"

func Bark() string {
	return "Woof!"
}

func  Barks() string {
  return "Woof! woof! woof!"
}

func BigBark() string {
  return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
  return dog.WhenGrownUp(Barks())
}
