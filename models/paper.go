package models

type Paper struct {
	BaseModel
	Title string `json:"title" gorm:"column:title"` // 试卷标题
}

func (m *Paper) TableName() string {
	return "paper"
}
