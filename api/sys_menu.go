package api

import (
	"demo/global/response"
	"demo/model/request"
	resp "demo/model/response"
	"demo/service"
	"fmt"
	"github.com/gin-gonic/gin"
)
//对应存放的方法
func GetMenu(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	err, menus := service.GetMenuTree(waitUse.AuthorityId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败，%v", err), c)
	} else {
		response.OkWithData(resp.SysMenusResponse{Menus: menus}, c)
	}
}
