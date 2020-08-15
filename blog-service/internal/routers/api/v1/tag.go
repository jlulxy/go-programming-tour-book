package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jlulxy/go-programming-tour-book/blog-service/global"
	"github.com/jlulxy/go-programming-tour-book/blog-service/internal/service"
	"github.com/jlulxy/go-programming-tour-book/blog-service/pkg/app"
	"github.com/jlulxy/go-programming-tour-book/blog-service/pkg/convert"
	"github.com/jlulxy/go-programming-tour-book/blog-service/pkg/errcode"
)

type Tag struct{}

func NewsTag() Tag {
	return Tag{}
}

func (a Tag) Get(c *gin.Context) {}

// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名称"
// @Success 200 {object} model.Tag "成功"
// @Router /api/v1/tags [get]
func (a Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	paper := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf(c, "svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	tags, err := svc.GetTagList(&param, &paper)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	response.ToResponseList(tags, totalRows)
	return
}
func (a Tag) Create(c *gin.Context) {
	params := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&params)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
func (a Tag) Update(c *gin.Context) {
	params := service.UpdateTagRequest{ID: convert.StrTo(c.Param("id")).MustUint32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&params)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return

}
func (a Tag) Delete(c *gin.Context) {
	params := service.DeleteTagRequest{ID: convert.StrTo(c.Param("id")).MustUint32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&params)
	if err != nil {
		global.Logger.Errorf(c, "svc.Delete err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
