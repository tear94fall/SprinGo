package domain

import (
	"main/dto"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Id     int
	Passwd string
	Email  string
}

func NewMember(memberDto dto.MemberDto) *Member {
	member := &Member{}
	member.Id = memberDto.GetId()
	member.Passwd = memberDto.GetPasswd()
	member.Email = memberDto.GetEmail()

	return member
}
