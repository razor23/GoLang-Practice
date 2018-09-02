package main

import "fmt"
import "math/rand"
import "time"

func main() {
  fmt. Println ("\n\t Welcome to V1N's Rock,Paper,Scissor \n")
	fmt.Println("Enter your play (R,P,S): ")
	var input string
	fmt.Scanln(&input)
	gameplay(input)
}

func gameplay(in string) {

	switch in {
	case "R":
		fmt.Println("You entered: 'Rock' ")
	case "P":
		fmt.Println("You entered: 'Paper' ")
	case "S":
		fmt.Println("You entered: 'Scissor' ")
	default:
		fmt.Println("Your entry is invalid")
	}

	a := random(0, 2)
	plays:= []string{"Rock","Paper","Scissor"}
  p:= plays[a]
	fmt.Println("\nComputer Played:",p)

switch p {

case "Rock":
  if in=="R"{fmt.Println("Go Again!")}
  if in=="P" {fmt.Println("You Win!")}
  if in=="S"{fmt.Println("You Lose!")}

case "Paper":
  if in=="R"{fmt.Println("You Lose!")}
  if in=="P" {fmt.Println("Go Again!")}
  if in=="S"{fmt.Println("You win!")}

case "Scissor":
  if in=="R"{fmt.Println("You Win!")}
  if in=="P" {fmt.Println("You Lose!")}
  if in=="S"{fmt.Println("Go Again!")}

default:
  (fmt.Println("This computer is tired of playing"))
}

}
func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)

}
