package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Question структура для хранения вопроса и ответа
type Question struct {
    question string
    answer   string
}

// generateQuestions генерирует список вопросов для теста
func generateQuestions() []Question {
    questions := make([]Question, 5)
    questions[0] = Question{question: "What is the capital of France?", answer: "Paris"}
    questions[1] = Question{question: "What is the largest planet in the solar system?", answer: "Jupiter"}
    questions[2] = Question{question: "Who wrote 'To Kill a Mockingbird'?", answer: "Harper Lee"}
    questions[3] = Question{question: "What is the square root of 64?", answer: "8"}
    questions[4] = Question{question: "Who is the current president of the United States?", answer: "Joe Biden"}
    return questions
}

// shuffleQuestions перемешивает вопросы в случайном порядке
func shuffleQuestions(questions []Question) {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })
}

// askQuestions задает вопросы пользователю и проверяет ответы
func askQuestions(questions []Question) int {
    correctAnswers := 0
    for i, question := range questions {
        fmt.Printf("%d. %s\n", i+1, question.question)
        var answer string
        fmt.Scanln(&answer)
        if answer == question.answer {
            correctAnswers++
        }
    }
    return correctAnswers
}

func main() {
    questions := generateQuestions()
    shuffleQuestions(questions)
    correctAnswers := askQuestions(questions)
    fmt.Printf("You got %d out of %d correct.\n", correctAnswers, len(questions))
}
