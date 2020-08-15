package service

import (
	"github.com/jlulxy/go-programming-tour-book/blog-service/internal/model"
	"github.com/jlulxy/go-programming-tour-book/blog-service/pkg/app"
)

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name     string `form:"name" binding:"required,min=1,max=100"`
	CreateBy string `form:"created_by" binding:"required,min=3,max=100"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountTag(params *CountTagRequest) (int, error) {
	return svc.dao.CountTag(params.Name, params.State)
}

func (svc *Service) GetTagList(params *TagListRequest, paper *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(params.Name, params.State, paper.Page, paper.PageSize)
}

func (svc *Service) CreateTag(params *CreateTagRequest) error {
	return svc.dao.CreateTag(params.Name, params.State, params.CreateBy)
}

func (svc *Service) UpdateTag(params *UpdateTagRequest) error {
	return svc.dao.UpdateTag(params.ID, params.Name, params.State, params.ModifiedBy)
}

func (svc *Service) DeleteTag(parmas *DeleteTagRequest) error {
	return svc.dao.DeleteTag(parmas.ID)
}
