package models

import "time"

type BaseModel struct {
	Id          int64     `json:"id,omitempty"`
	GmtCreated  time.Time `json:"gmt_created"`
	GmtModified time.Time `json:"gmt_modified"`
	IsDeleted   int       `json:"is_deleted,omitempty"`
}
