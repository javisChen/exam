package models

type UserQuestionRecord struct {
	BaseModel
	QuestionID uint64 `json:"question_id" gorm:"column:question_id"` // 试题id,关联question.id
	PaperID    uint64 `json:"paper_id" gorm:"column:paper_id"`       // 试卷id,关联paper.id
	IsRight    int8   `json:"is_right" gorm:"column:is_right"`       // 是否正确 0-否 1-是
	UserAnswer string `json:"user_answer" gorm:"column:user_answer"` // 试题答案，单选/多选题存储选项的code，以,分割，例如1,2,3
	Status     int8   `json:"status" gorm:"column:status"`           // 答题状态 0-未完成 1-已完成
}

func (m *UserQuestionRecord) TableName() string {
	return "user_question_record"
}
