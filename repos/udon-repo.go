package repos

import (
	"fmt"
	"udon-web/models"
)

type UdonRepo struct {
	udons []models.Udon
}

func NewUdonRepo() *UdonRepo {
	return &UdonRepo{make([]models.Udon, 0)}
}

func (p *UdonRepo) Create(partial models.Udon) models.Udon {
	newItem := partial
	newItem.ID = uint(len(p.udons)) + 1
	p.udons = append(p.udons, newItem)
	return newItem
}

func (p *UdonRepo) GetList() []models.Udon {
	return p.udons
}

func (p *UdonRepo) GetOne(id uint) (models.Udon, error) {
	for _, it := range p.udons {
		if it.ID == id {
			return it, nil
		}
	}
	return models.Udon{}, fmt.Errorf("key '%d' not found", id)
}

func (p *UdonRepo) Update(id uint, amended models.Udon) (models.Udon, error) {
	for i, it := range p.udons {
		if it.ID == id {
			amended.ID = id
			p.udons = append(p.udons[:i], p.udons[i+1:]...)
			p.udons = append(p.udons, amended)
			return amended, nil
		}
	}
	return models.Udon{}, fmt.Errorf("key '%d' not found", amended.ID)
}

func (p *UdonRepo) DeleteOne(id uint) (bool, error) {
	for i, it := range p.udons {
		if it.ID == id {
			p.udons = append(p.udons[:i], p.udons[i+1:]...)
			return true, nil
		}
	}
	return false, fmt.Errorf("key '%d' not found", id)
}