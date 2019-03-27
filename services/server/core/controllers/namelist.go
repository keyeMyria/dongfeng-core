package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// NamelistController Namelist controller
type NamelistController struct {
	repository *repositories.NamelistRepository
}

// NewNamelistController new controller
func NewNamelistController() *NamelistController {
	return &NamelistController{
		repository: repositories.NewNamelistRepository(),
	}
}

// GetNamelists get namelists
func (c *NamelistController) GetNamelists(class, year string) ([]*models.Namelist, error) {
	return c.repository.Select(class, year)
}

// UpdateNamelists update namelists
func (c *NamelistController) UpdateNamelists(items []*proto.NamelistItem) error {
	return c.repository.DeleteInsert(items)
}
