package live_chat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	lcModel "github.com/flipped-aurora/gin-vue-admin/server/model/live_chat"

	"github.com/gin-gonic/gin"
)

type ProductApi struct{}

func (a *ProductApi) Create(c *gin.Context) {
	var p lcModel.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lcServiceGroup.ProductService.Create(&p); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (a *ProductApi) Update(c *gin.Context) {
	var p lcModel.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lcServiceGroup.ProductService.Update(&p); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (a *ProductApi) Delete(c *gin.Context) {
	var req request.GetById
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lcServiceGroup.ProductService.Delete(int64(req.ID)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (a *ProductApi) List(c *gin.Context) {
	var search lcModel.ProductSearch
	var pageInfo request.PageInfo
	if err := c.ShouldBindQuery(&search); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := lcServiceGroup.ProductService.List(search, pageInfo.Page, pageInfo.PageSize)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{List: list, Total: total, Page: pageInfo.Page, PageSize: pageInfo.PageSize}, "获取成功", c)
}

func (a *ProductApi) Detail(c *gin.Context) {
	id := c.Param("id")
	var req request.GetById
	req.ID = 0
	if id != "" {
		c.ShouldBindUri(&req)
	}
	p, err := lcServiceGroup.ProductService.GetByID(int64(req.ID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(p, c)
}
