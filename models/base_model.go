package models

import "time"

type BaseModel struct {
	Id          uint64    `json:"id,omitempty"`
	GmtCreated  time.Time `json:"gmt_created" orm:"-"`
	GmtModified time.Time `json:"gmt_modified" orm:"-"`
	IsDeleted   uint      `json:"is_deleted,omitempty" orm:"-"`
}
