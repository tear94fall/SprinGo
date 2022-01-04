package controller

import (
	"main/dto"
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
	MemberService *service.MemberService
}

func NewMemberController() *MemberController {
	memberController := &MemberController{}
	memberController.MemberService = service.NewMemberService()

	return memberController
}

func (memberController MemberController) Register(c *gin.Context) {
	member := dto.MemberDto{}

	err := c.ShouldBind(&member)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad reqeust",
		})

		return
	}

	if err := memberController.MemberService.Register(member); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "Register success",
	})
}

func (memberController MemberController) List(c *gin.Context) {
	memberList := memberController.MemberService.List()

	c.JSON(200, gin.H{
		"message":    "List success",
		"memberList": memberList,
	})
}
