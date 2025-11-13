package handler

import (
	"github.com/gin-gonic/gin"
)

type QuestionService interface {
}

type QuestionHandler struct {
	QuestionService QuestionService
}

func NewQuestionHandler(questionService QuestionService) *QuestionHandler {
	return &QuestionHandler{QuestionService: questionService}
}

func (h *QuestionHandler) GetQuestions(c *gin.Context) {
	// var questions []entity.Question
	// result := database.DB.Find(&questions)
	// if result.Error != nil {
	// 	http.Error(w, "Error fetching questions", http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(questions)
}

// CreateQuestion создает новый вопрос
func (h *QuestionHandler) CreateQuestion(c *gin.Context) {
	// var question entity.Question
	// if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
	// 	http.Error(w, "Invalid request body", http.StatusBadRequest)
	// 	return
	// }

	// if question.Text == "" {
	// 	http.Error(w, "Question text is required", http.StatusBadRequest)
	// 	return
	// }

	// question.CreatedAt = time.Now()

	// result := database.DB.Create(&question)
	// if result.Error != nil {
	// 	http.Error(w, "Error creating question", http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(question)
}

// GetQuestion возвращает вопрос и все ответы на него
func (h *QuestionHandler) GetQuestion(c *gin.Context) {
	// vars := mux.Vars(r)
	// id, err := strconv.Atoi(vars["id"])
	// if err != nil {
	// 	http.Error(w, "Invalid question ID", http.StatusBadRequest)
	// 	return
	// }

	// var question entity.Question
	// result := database.DB.Preload("Answers").First(&question, id)
	// if result.Error != nil {
	// 	http.Error(w, "Question not found", http.StatusNotFound)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(question)
}

// DeleteQuestion удаляет вопрос и все его ответы (каскадно)
func (h *QuestionHandler) DeleteQuestion(c *gin.Context) {
	// vars := mux.Vars(r)
	// id, err := strconv.Atoi(vars["id"])
	// if err != nil {
	// 	http.Error(w, "Invalid question ID", http.StatusBadRequest)
	// 	return
	// }

	// result := database.DB.Delete(&entity.Question{}, id)
	// if result.Error != nil {
	// 	http.Error(w, "Error deleting question", http.StatusInternalServerError)
	// 	return
	// }

	// if result.RowsAffected == 0 {
	// 	http.Error(w, "Question not found", http.StatusNotFound)
	// 	return
	// }

	// w.WriteHeader(http.StatusNoContent)
}
