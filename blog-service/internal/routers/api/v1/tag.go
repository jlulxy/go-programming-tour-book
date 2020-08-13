package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jlulxy/go-programming-tour-book/blog-service/global"
	"github.com/jlulxy/go-programming-tour-book/blog-service/pkg/app"
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
	params := struct {
		Name  string `form:"name" binding:"max=100"`
		State uint8  `form:"state,default=1" binding:"oneof=0 1"`
	}{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	response.ToResponse(gin.H{})
	return
}
func (a Tag) Create(c *gin.Context) {}
func (a Tag) Update(c *gin.Context) {}
func (a Tag) Delete(c *gin.Context) {}
