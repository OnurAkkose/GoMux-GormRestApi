package service

import (
	model "gorestapi/pkg/model/member"
	member "gorestapi/pkg/repository/member"
)

type MemberService struct {
	MemberRepository *member.MemberRepository
}

func NewMemberService(m *member.MemberRepository) MemberService {
	return MemberService{MemberRepository: m}
}

func (m *MemberService) All() ([]model.Member, error) {
	return m.MemberRepository.All()
}

func (m *MemberService) FindById(id uint) (*model.Member, error) {
	return m.MemberRepository.FindById(id)
}

func (m *MemberService) Save(member *model.Member) (*model.Member, error) {
	return m.MemberRepository.Save(member)
}

func (m *MemberService) Delete(id uint) error {
	return m.MemberRepository.Delete(id)
}

func (m *MemberService) Migrate() error {
	return m.MemberRepository.Migrate()
}
