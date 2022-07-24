package comment

import (
	"github.com/motikingo/resturant-api/entity"
)
type CommentService interface{
	Comments() ([]entity.Comment,[]error)
	Comment(id uint) (*entity.Comment,[]error)
	UpdateComment(comm *entity.Comment)(*entity.Comment,[]error)
	DeleteComment(id uint)(*entity.Comment,[]error)
	CreateComment(comm *entity.Comment)(*entity.Comment,[]error)
}