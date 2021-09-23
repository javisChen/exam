package models

import "time"

type BaseModel struct {
	Id          int64     `json:"id,omitempty" orm:"column(id);"`
	GmtCreated  time.Time `json:"gmt_created" orm:"column(gmt_created);type(datetime)"`
	GmtModified time.Time `json:"gmt_modified" orm:"column(gmt_modified);type(datetime)"`
	IsDeleted   int       `json:"is_deleted,omitempty" orm:"column(is_deleted);"`
}
