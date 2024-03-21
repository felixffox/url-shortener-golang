package storage

import "errors"

var (
	ErrURLNotFound = errors.New("URL not found")
	ErrURLExist    = errors.New("URL exist")
)

type URLInfo struct {
	ID    int64  `json:"id" gorm:"primaryKey; not null"`
	URL   string `json:"url"`
	Alias string `json:"alias"`
}

type URLSaver interface {
	SaveURL(urlToSave string, alias string) (int64, error)
}

type URLLister interface {
	GetAllURLs() ([]URLInfo, error)
}

type URLDeleter interface {
	DeleteURL(alias string) error
}
