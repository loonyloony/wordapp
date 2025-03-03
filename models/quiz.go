package models

type QuizAnswerEvent struct {
	QuizID     string `json:"quiz_id"`
	UserID     string `json:"user_id"`
	QuestionID string `json:"question_id"`
	Answer     string `json:"answer"`
	Type       string `json:"type"`
	Timestamp  int64  `json:"timestamp"`
}

type QuizAnswerResult struct {
	Score   int `json:"score"`
	Correct int `json:"correct"`
}
