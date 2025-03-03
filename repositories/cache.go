package repositories

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type QuizScoreRepository interface {
	UpdateScore(quizID, userID string, score int) error
}

type QuizScore struct {
	client *redis.Client
}

func NewQuizScoreRepository(client *redis.Client) QuizScoreRepository {
	return &QuizScore{
		client: client,
	}
}

func (r *QuizScore) UpdateScore(quizID, userID string, score int) error {
	return r.client.ZIncrBy(context.Background(), quizID+":scores", float64(score), userID).Err()
}
