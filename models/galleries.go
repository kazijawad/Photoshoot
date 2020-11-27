package models

import "github.com/jinzhu/gorm"

// Gallery represents the galleries table in our DB
// and is mostly a container resource composed of images.
type Gallery struct {
	gorm.Model
	UserID uint `gorm:"not_null;index"`
	Title  uint `gorm:"not_null"`
}

type GalleryService interface {
	GalleryDB
}

type GalleryDB interface {
	Create(gallery *Gallery) error
}

type galleryGorm struct {
	db *gorm.DB
}

func (gg *galleryGorm) Create(gallery *Gallery) error {
	// TODO: Implement
	return nil
}
