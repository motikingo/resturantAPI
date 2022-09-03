package CommentService

import (
	"github.com/motikingo/resturant-api/comment"

	"github.com/motikingo/resturant-api/entity"
)

type Commentsrv struct {
	repo comment.CommentRepository
}

func NewCommentService(repo comment.CommentRepository) comment.CommentService {

	return &Commentsrv{repo: repo}
}

func (comsrv *Commentsrv) Comments() ([]entity.Comment, []error) {
	comments, err := comsrv.repo.Comments()

	if len(err) > 0 {
		return nil, err
	}

	return comments, nil
}

func (comsrv *Commentsrv) Comment(id uint) (*entity.Comment, []error) {
	cmt, err := comsrv.repo.Comment(id)
	if len(err) > 0 {
		return nil, err
	}
	return cmt, nil

}

func (comsrv *Commentsrv) UpdateComment(comm entity.Comment) (*entity.Comment, []error) {
	cmt, err := comsrv.repo.UpdateComment(comm)

	if len(err) > 0 {
		return nil, err
	}
	return cmt, nil
}

func (comsrv *Commentsrv) DeleteComment(id uint) (*entity.Comment, []error) {
	cmt, err := comsrv.repo.DeleteComment(id)

	if len(err) > 0 {
		return nil, err

	}

	return cmt, nil
}

func (comsrv *Commentsrv) CreateComment(cm *entity.Comment) (*entity.Comment, []error) {
	comment, err := comsrv.repo.CreateComment(cm)
	if len(err) > 0 {
		return nil, err
	}
	return comment, err
}
