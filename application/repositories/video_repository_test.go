package repositories_test

import (
	"video-encoder/application/repositories"
	"video-encoder/domain"
	"testing"
	"github.com/stretchr/testify/require"
	"time"
	uuid "github.com/satori/go.uuid"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.NewVideoRepository(db)
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, video.ID)
}
