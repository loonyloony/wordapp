package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"wordapp/models"
	"wordapp/repositories"
	"wordapp/services"

	"github.com/redis/go-redis/v9"
)

func main() {
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost:6379"
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisHost,
	})

	quizScoreRepo := repositories.NewQuizScoreRepository(redisClient)
	quizScoreServ := services.NewQuizScoreService(quizScoreRepo)

	// HTTP endpoint để nhận câu trả lời từ Postman
	http.HandleFunc("/quiz/answer", func(w http.ResponseWriter, r *http.Request) {
		var answer models.QuizAnswerEvent
		if err := json.NewDecoder(r.Body).Decode(&answer); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		err := quizScoreServ.ProcessAnswer(&answer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "Answer processed"})
	})

	log.Println("Score Service started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
