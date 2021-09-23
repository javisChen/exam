package req

type PaperCreateReq struct {
	Questions []Questions `json:"questions"`
	Title     string      `json:"title"`
}

type Questions struct {
	Answer  int              `json:"answer"`
	Options []QuestionOption `json:"options"`
	Title   string           `json:"title"`
	Type    int              `json:"type"`
}

type QuestionOption struct {
	Title string `json:"title"`
	Seq   int    `json:"seq"`
}
