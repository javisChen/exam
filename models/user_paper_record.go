package models

type UserPagerRecord struct {
	BaseModel
	PaperID uint64 `json:"paper_id" gorm:"column:paper_id"` // 试卷id,关联paper.id
	Score   int    `json:"score" gorm:"column:score"`       // 获得的分数
	Status  int8   `json:"status" gorm:"column:status"`     // 状态 0-未完成 1-已完成
}

func (m *UserPagerRecord) TableName() string {
	return "user_pager_record"
}
