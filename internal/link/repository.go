package link

import (
	"golang-app/pkg/db"
	"log"
)

type LinkRepository struct {
	Database *db.DB
}

func NewLinkRepository(db *db.DB) *LinkRepository {
	return &LinkRepository{
		Database: db,
	}
}

func (repository *LinkRepository) Create(link *Link) error {

	res := repository.Database.Create(link)

	if res.Error != nil {
		return res.Error
	}

	log.Print("Link created: ", link)

	return nil
}

func (repository *LinkRepository) Edit(link *Link) error {
	return nil
}

func (repository *LinkRepository) Delete(id string) error {

	res := repository.Database.Delete(&Link{}, id)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
