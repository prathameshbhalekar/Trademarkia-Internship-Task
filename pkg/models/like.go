package models

type Like struct {
	ID           uint `json:"id" gorm:"primary_key;not null"`
	WhoLikesID   uint `json:"who_likes" gorm:"not null"`
	WhoLikes     Users
	WhoIsLikedID uint `json:"who_is_liked" gorm:"not null"`
	WhoIsLiked   Users
}
