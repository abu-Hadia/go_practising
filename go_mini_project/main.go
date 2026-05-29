package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Question struct {
	Text   string
	Answer string
}

var questions = map[int]Question{
	1: {"Sanadkee ayuu nabigu dhashay?", "571"},
	2: {"Halkuu ku dhashay?", "Makka"},
	3: {"Imisa xaas ayuu lahaa?", "11"},
	4: {"Imisa khulafaa raashidiin?", "4"},
}

var score = 0
var currentQ int = 0

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n🎮 Quiz Game")
		fmt.Println("Pick a number (1-4) or 0 to exit:")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("❌ Invalid input")
			continue
		}

		if num == 0 {
			fmt.Println("🏁 Game Over!")
			fmt.Println("Final Score:", score)
			break
		}

		q, exists := questions[num]
		if !exists {
			fmt.Println("⚠️ No question for that number")
			continue
		}

		currentQ = num

		fmt.Println("\n❓ Question:", q.Text)
		fmt.Println("⏱ You have 10 seconds to answer...")

		answerChan := make(chan string)

		// goroutine for input
		go func() {
			ans, _ := reader.ReadString('\n')
			answerChan <- strings.TrimSpace(ans)
		}()

		select {
		case ans := <-answerChan:
			if strings.EqualFold(ans, q.Answer) {
				fmt.Println("✅ Correct!")
				score++
			} else {
				fmt.Println("❌ Wrong! Correct answer:", q.Answer)
			}

		case <-time.After(10 * time.Second):
			fmt.Println("⏱ Time's up!")
		}

		fmt.Println("📊 Score:", score)
	}
}
