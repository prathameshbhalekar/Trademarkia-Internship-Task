package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prathameshbhalekar/trademarkia-task/pkg/db"
)

type Matches struct {
	Match_User1_ID int
	Match_User2_ID int
}

func GetAllMatches(c *gin.Context) {

	query := `SELECT A.ID AS Match_User1_ID, B.ID AS Match_User2_ID 
		FROM Likes AS A
		INNER JOIN 
		Likes AS B
		ON A.who_likes_id = B.who_is_liked_id
		WHERE A.who_likes_id = B.who_is_liked_id
		AND B.who_likes_id = A.who_is_liked_id
		AND a.who_likes_id < b.who_likes_id
	`

	var matches []Matches
	db.DB.Raw(query).Scan(&matches)

	c.JSON(http.StatusOK, gin.H{"data": matches})

}
