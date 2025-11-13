package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	QuestionHandler *QuestionHandler
	AnswerHandler   *AnswerHandler
}

func NewHandler(s ser.Service) *Handler {
	return &Handler{
		QuestionHandler: NewQuestionHandler(s),
		AnswerHandler:   NewAnswerHandler(s),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api/v1")
	{
		questions := api.Group("/questions")
		{
			questions.GET("/", h.QuestionHandler.GetQuestion)
			questions.POST("/", h.QuestionHandler.CreateQuestion)
			questions.GET("/:id", h.QuestionHandler.GetQuestion)
			questions.DELETE("/:id", h.QuestionHandler.DeleteQuestion)
		}

		answers := api.Group("/answers")
		{
			answers.POST("/questions/:id/answers", h.AnswerHandler.DeleteAnswer)
			answers.GET("/:id", h.AnswerHandler.GetAnswer)
			answers.DELETE("/:id", h.AnswerHandler.DeleteAnswer)

		}
	}

	return router
}
