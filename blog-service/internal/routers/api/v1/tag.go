package v1

import "github.com/gin-gonic/gin"

type Tag struct{}

func NewsTag() Tag {
	return Tag{}
}

func (a Tag) Get(c *gin.Context)    {}
func (a Tag) List(c *gin.Context)   {}
func (a Tag) Create(c *gin.Context) {}
func (a Tag) Update(c *gin.Context) {}
func (a Tag) Delete(c *gin.Context) {}
