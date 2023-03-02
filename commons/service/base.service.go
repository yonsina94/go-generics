package service

import (
	"gorm.io/gorm"
)

type BaseService[Tmodel any] struct {
	DB *gorm.DB
}

func (s *BaseService[Tmodel]) GetAll() (data []Tmodel, err error) {
	err = s.DB.Model(new(Tmodel)).Find(&data).Error

	return
}

func (s *BaseService[Tmodel]) GetBy(conditions Tmodel) (data []Tmodel, err error) {
	err = s.DB.Model(new(Tmodel)).Find(&data, &conditions).Error

	return
}

func (s *BaseService[Tmodel]) GetByID(id uint) (data Tmodel, err error) {
	err = s.DB.Model(new(Tmodel)).Find(&data, map[string]any{"ID": id}).Error

	return
}

func (s *BaseService[Tmodel]) Save(data ...Tmodel) (bool, error) {
	if err := s.DB.Model(new(Tmodel)).Save(&data).Error; err == nil {
		return true, nil
	} else {
		return false, err
	}
}

func (s *BaseService[Tmodel]) DeleteById(id uint) (bool, error) {
	if err := s.DB.Delete(new(Tmodel), id).Error; err == nil {
		return true, nil
	} else {
		return false, err
	}
}
