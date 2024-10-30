package coba

import (
	"anograph/graph/model"
	"gorm.io/gorm"
)
type CobaRepository interface {
	GetAll() ([]*model.Coba, error)
	GetByID(id string) (*model.Coba, error)
	Create(coba *model.Coba) (*model.Coba, error)
	Update(coba *model.Coba) (*model.Coba, error)
	Delete(id string) error
}

type CobaRepositoryImpl struct {
	db *gorm.DB
}

func NewCobaRepository(db *gorm.DB) CobaRepository {
	return &CobaRepositoryImpl{db}
}

func (r *CobaRepositoryImpl) GetAll() ([]*model.Coba, error) {
	var cobas []*model.Coba
	err := r.db.Find(&cobas).Error
	return cobas, err
}

func (r *CobaRepositoryImpl) GetByID(id string) (*model.Coba, error) {
	var coba model.Coba
	err := r.db.First(&coba, "id = ?", id).Error
	return &coba, err
}

func (r *CobaRepositoryImpl) Create(coba *model.Coba) (*model.Coba, error) {
	err := r.db.Create(coba).Error
	return coba, err
}

func (r *CobaRepositoryImpl) Update(coba *model.Coba) (*model.Coba, error) {
	err := r.db.Save(coba).Error
	return coba, err
}

func (r *CobaRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&model.Coba{}, "id = ?", id).Error
}
