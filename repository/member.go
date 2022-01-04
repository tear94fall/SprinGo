package repository

import (
	"main/database"
	"main/domain"
	"main/dto"
)

type MemberRepository struct {
	EntityManger *database.Database
}

func NewMemberRepository() *MemberRepository {
	MemberRepository := &MemberRepository{}
	MemberRepository.EntityManger = database.GetDbInstance()

	return MemberRepository
}

func (memberRepository MemberRepository) Create(memberDto dto.MemberDto) error {
	member := domain.NewMember(memberDto)

	err := memberRepository.EntityManger.Save(member)
	if err != nil {
		return err
	}

	return nil
}

func (memberRepository MemberRepository) List() []domain.Member {
	var memberList []domain.Member

	memberRepository.EntityManger.FindAll(&memberList)

	return memberList
}
