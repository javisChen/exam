package resp

type PaperInfoResp struct {
	Id        int64       `json:"id"`
	Questions []Questions `json:"questions"`
	Title     string      `json:"title"`
}

type Questions struct {
	Id      int64            `json:"id"`
	Answer  string           `json:"answer"`
	Options []QuestionOption `json:"options"`
	Title   string           `json:"title"`
	Type    int              `json:"type"`
}

type QuestionOption struct {
	Id            int64    `json:"id"`
	Title         string   `json:"title"`
	Seq           int      `json:"seq"`
	CheckedValues []string `json:"checkedValues"`
}
