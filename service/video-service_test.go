package service

import (
	"testing"

	"github.com/caohieu04/Gin-Course/entity"
	"github.com/stretchr/testify/assert"
)

const (
	TITLE       = "Video Title"
	DESCRIPTION = "Video Description"
	URL         = "https://youtu.be"
)

func getVideo() entity.Video {
	return entity.Video{
		Title:       TITLE,
		Description: DESCRIPTION,
		URL:         URL,
	}
}

func TestFindAll(t *testing.T) {
	service := New()

	service.Save(getVideo())

	videos := service.FindAll()

	firstVideo := videos[0]
	assert.Nil(t, videos)
	assert.Equal(t, TITLE, firstVideo.Title)
	assert.Equal(t, DESCRIPTION, firstVideo.Description)
	assert.Equal(t, URL, firstVideo.URL)
}
