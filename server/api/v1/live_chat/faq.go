package live_chat

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	lcModel "github.com/flipped-aurora/gin-vue-admin/server/model/live_chat"

	"github.com/gin-gonic/gin"
)

type FaqApi struct{}

func (a *FaqApi) Create(c *gin.Context) {
	var f lcModel.Faq
	if err := c.ShouldBindJSON(&f); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lcServiceGroup.FaqService.Create(&f); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (a *FaqApi) Update(c *gin.Context) {
	var f lcModel.Faq
	if err := c.ShouldBindJSON(&f); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lcServiceGroup.FaqService.Update(&f); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (a *FaqApi) Delete(c *gin.Context) {
	var req request.GetById
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lcServiceGroup.FaqService.Delete(int64(req.ID)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (a *FaqApi) List(c *gin.Context) {
	var search lcModel.FaqSearch
	var pageInfo request.PageInfo
	if err := c.ShouldBindQuery(&search); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := lcServiceGroup.FaqService.List(search, pageInfo.Page, pageInfo.PageSize)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{List: list, Total: total, Page: pageInfo.Page, PageSize: pageInfo.PageSize}, "获取成功", c)
}

func (a *FaqApi) Import(c *gin.Context) {
	var req lcModel.FaqImportReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	for i := range req.Items {
		req.Items[i].ProductID = req.ProductID
		lcServiceGroup.FaqService.Create(&req.Items[i])
	}
	response.OkWithMessage("导入成功", c)
}

func (a *FaqApi) Categories(c *gin.Context) {
	productID := int64(0)
	if v, ok := c.GetQuery("product_id"); ok && v != "" {
		var pid int
		if _, err := fmt.Sscanf(v, "%d", &pid); err == nil {
			productID = int64(pid)
		}
	}
	cats, err := lcServiceGroup.FaqService.Categories(productID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(cats, c)
}
