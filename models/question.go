package models

type Question struct {
	BaseModel
	PaperID uint64 `json:"paper_id" gorm:"column:paper_id"` // 试卷id,关联paper.id
	Title   string `json:"title" gorm:"column:title"`       // 试题标题
	Type    int8   `json:"type" gorm:"column:type"`         // 试题类型 1-单选 2-多选 3-问答
	Score   int    `json:"score" gorm:"column:score"`       // 试题分数
	Answer  string `json:"answer" gorm:"column:answer"`     // 试题答案，问答题答案为空，单选/多选题存储选项的code，以,分割，例如1,2,3
}

func (m *Question) TableName() string {
	return "question"
}
