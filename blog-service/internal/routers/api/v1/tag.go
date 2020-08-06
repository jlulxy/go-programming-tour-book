package v1

import "github.com/gin-gonic/gin"

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
func (a Tag) List(c *gin.Context)   {}
func (a Tag) Create(c *gin.Context) {}
func (a Tag) Update(c *gin.Context) {}
func (a Tag) Delete(c *gin.Context) {}
