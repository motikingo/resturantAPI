package CommentRepository

import (
	"github.com/jinzhu/gorm"
	"github.com/motikingo/resturant-api/comment"
	"github.com/motikingo/resturant-api/entity"
)

type CommentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) comment.CommentRepository {

	return &CommentRepo{db: db}
}

func (com *CommentRepo) Comments() ([]entity.Comment, []error) {
	comments := []entity.Comment{}

	err := com.db.Find(&comments).GetErrors()

	if len(err) > 0 {
		return nil, err
	}

	return comments, nil
}

func (com *CommentRepo) Comment(id uint) (*entity.Comment, []error) {
	var cmt entity.Comment

	err := com.db.First(&cmt, id).GetErrors()
	if len(err) > 0 {
		return nil, err
	}
	return &cmt, nil

}

func (com *CommentRepo) UpdateComment(comm entity.Comment) (*entity.Comment, []error) {
	cmt, er := com.Comment(comm.ID)
	if len(er) > 0 {
		return nil, er
	}
	cmt.Description = func() string {
		if comm.Description != cmt.Description {
			return comm.Description
		}
		return cmt.Description
	}()

	if err := com.db.Save(cmt).GetErrors(); len(err) > 0 {
		return nil, err
	}
	return cmt, nil
}

func (com *CommentRepo) DeleteComment(id uint) (*entity.Comment, []error) {
	cmt, err := com.Comment(id)

	if len(err) > 0 {
		return nil, err

	}
	err = com.db.Delete(&cmt, id).GetErrors()
	if len(err) > 0 {
		return nil, err

	}

	return cmt, nil
}

func (com *CommentRepo) CreateComment(cm *entity.Comment) (*entity.Comment, []error) {
	comment := cm
	err := com.db.Create(comment).GetErrors()
	if len(err) > 0 {
		return nil, err
	}
	return comment, nil
}
