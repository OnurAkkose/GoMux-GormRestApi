package repository

import (
	model "gorestapi/pkg/model/member"

	"github.com/jinzhu/gorm"
)

type MemberRepository struct {
	db *gorm.DB
}

func NewMember(db *gorm.DB) *MemberRepository {
	return &MemberRepository{
		db: db,
	}
}

func (m *MemberRepository) All() ([]model.Member, error) {
	members := []model.Member{}
	err := m.db.Find(&members).Error
	return members, err
}

func (m *MemberRepository) FindById(id uint) (*model.Member, error) {
	member := new(model.Member)
	err := m.db.Where(`id = ?`, id).First(&member).Error
	return member, err
}

func (m *MemberRepository) Save(member *model.Member) (*model.Member, error) {
	err := m.db.Save(&member).Error
	return member, err
}

func (m *MemberRepository) Delete(id uint) error {
	err := m.db.Delete(&model.Member{ID: id}).Error
	return err
}

func (m *MemberRepository) Migrate() error {
	return m.db.AutoMigrate(&model.Member{}).Error

}
