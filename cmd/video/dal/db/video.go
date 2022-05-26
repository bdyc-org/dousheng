package db

import (
	"github.com/bdyc-org/dousheng/pkg/constants"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Title          string `json:"title"`
	Play_url       string `json:"play_url"`
	Cover_url      string `json:"cover_url"`
	User_id        uint   `json:"user_id"`
	Favorite_count int    `json:"favorite_count"`
	Comment_count  int    `json:"comment_count"`
}

func (v *Video) TableName() string{
	return constants.VideoTableName
}

