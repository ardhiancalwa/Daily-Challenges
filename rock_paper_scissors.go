package kata


func Rps(p1, p2 string) string {
  if (p1 == "scissors" && p2 == "paper") || (p1 == "paper" && p2 == "rock") || (p1 == "rock" && p2 == "scissors") {
    return "Player 1 won!"
  } else if (p2 == "scissors" && p1 == "paper") || (p2 == "paper" && p1 == "rock") || (p2 == "rock" && p1 == "scissors") {
    return "Player 2 won!"
  } else {
    return "Draw!"
  }
}
