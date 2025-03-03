package services

import (
	"log"

	"wordapp/models"
	"wordapp/repositories"
)

type QuizScoreService struct {
	quizScoreRepo repositories.QuizScoreRepository
}

func NewQuizScoreService(quizScoreRepo repositories.QuizScoreRepository) *QuizScoreService {
	return &QuizScoreService{
		quizScoreRepo: quizScoreRepo,
	}
}

func (s *QuizScoreService) ProcessAnswer(event *models.QuizAnswerEvent) error {
	answerResult, err := s.calculateScore(event)
	if err != nil {
		log.Printf("Error answer score calculating: %v", err)
		return err
	}

	err = s.quizScoreRepo.UpdateScore(event.QuizID, event.UserID, answerResult.Score)
	if err != nil {
		log.Printf("Error showing leaderboard: %v", err)
		return err
	}

	return nil
}

func (s *QuizScoreService) calculateScore(event *models.QuizAnswerEvent) (models.QuizAnswerResult, error) {
	// call api to quiz service to check answer
	// cache to enhance performance
	switch event.Type {
	case "multiple_choice":
		return models.QuizAnswerResult{Score: 20, Correct: 1}, nil
	default:
		return models.QuizAnswerResult{Score: 0, Correct: 0}, nil
	}
}
