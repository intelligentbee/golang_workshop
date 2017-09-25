package main

import(
    "fmt"
    "math/rand"
    "time"
)

func generateNumber(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main() {
    const Max = 10
    myrand := generateNumber(1, Max)
    guessTaken := 0

    var guess int

    fmt.Printf("Please guess the number between 1-10:\n")

    for {
        _, err := fmt.Scanf("%d", &guess)
        guessTaken++

        if err != nil {
          fmt.Println(err)
          continue
        }

        if guess != myrand {
            fmt.Printf("It isn’t %d. Try again.\n", guess)
        }

        if guess > Max {
            fmt.Printf("The max number is %d. Try again.\n", Max)
        }

        if guess == myrand {
            break
        }
    }

    if guess == myrand {
      if guessTaken > 1 {
        fmt.Printf("Yay! That’s the right number. It took you %d tries!\n", guessTaken)
      } else {
        fmt.Printf("Yay! That’s the right number. It took you %d try!\n", guessTaken)
      }
    } else {
        fmt.Printf("It isn’t %d. Try again.\n", myrand)
    }
}
