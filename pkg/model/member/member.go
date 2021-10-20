package model

import "time"

type Member struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	FullName  string     `json:"fullname"`
	Email     string     `json:"email"`
}

type MemberDto struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

func ToMember(memberDto MemberDto) *Member { // converts authordto to author
	return &Member{
		FullName: memberDto.FullName,
		Email:    memberDto.Email,
	}
}

func ToMemberDto(member *Member) *MemberDto { // converts author to authorDto
	return &MemberDto{
		ID:       member.ID,
		FullName: member.FullName,
		Email:    member.Email,
	}
}
