package service

import (
	"main/domain"
	"main/dto"
	"main/repository"
)

type MemberService struct {
	MemberRepository *repository.MemberRepository
}

func NewMemberService() *MemberService {
	memberService := &MemberService{}
	memberService.MemberRepository = repository.NewMemberRepository()

	return memberService
}

func (memberService MemberService) Register(member dto.MemberDto) error {
	err := memberService.MemberRepository.Create(member)
	if err != nil {
		return err
	}
	return nil
}

func (MemberService MemberService) List() []domain.Member {
	return MemberService.MemberRepository.List()
}
