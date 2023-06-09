package repository

import (
	"go-mongo-tutorial/src/modules/profile/model"
)

// profileRepository interface
type ProfileRepository interface {
	Save(*model.Profile) error
	Update(string, *model.Profile) error
	Delete(string) error
	FindByID(string) (*model.Profile, error)
	FindAll() (model.Profiles, error)
}
