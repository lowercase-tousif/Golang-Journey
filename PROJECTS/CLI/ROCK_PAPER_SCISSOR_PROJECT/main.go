package main

import (
	"fmt"
)

func main(){
	fmt.Println("Welcome to Rock Paper Scissor Game")
	fmt.Println("Game made by Tousif Tasrik")

}	


/*
   BREAKDOWN OF THE WINNING AND LOSING SCENARIOS

   WIN SEQUENCES

   Rock --> Lizard, Scissors
   Paper --> Rock, Spock
   Scissors --> Paper, Lizard
   Lizard --> Spock, Paper
   Spock --> Scissors, Rock

   LOSE SEQUENCES

   Rock --> Paper, Spock
   Paper --> Scissors, Lizard
   Scissors --> Rock, Spock
   Lizard --> Rock, Scissors
   Spock --> Lizard, Paper
*/
