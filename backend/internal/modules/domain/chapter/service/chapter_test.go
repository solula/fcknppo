package service

import (
	"context"
	"fmt"
	"testing"
	"waterfall-backend/internal/modules/domain/chapter/dto"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type txMock struct {
}

func (t *txMock) Commit() error {
	return nil
}

func (t *txMock) Rollback() error {
	return nil
}

// TestChapterService_GetByUuid
func TestChapterService_GetByUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIChapterRepo(ctrl)
	bc := NewMockIBucketCreator(ctrl)
	service := NewChapterService(repo, bc)

	uuid := "123e4567-e89b-12d3-a456-426655440000"
	expectedChapter := &dto.Chapter{
		Uuid:        "123e4567-e89b-12d3-a456-426655440000",
		Number:      1,
		Title:       "Example Chapter",
		PartUuid:    "part-uuid",
		ReleaseUuid: nil,
	}

	repo.EXPECT().GetByUuid(gomock.Any(), uuid).Return(expectedChapter, nil)

	chapter, err := service.GetByUuid(context.Background(), uuid)
	if assert.NoError(t, err) {
		assert.Equal(t, expectedChapter, chapter)
	}
}

func TestChapterService_GetByUuid_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIChapterRepo(ctrl)
	bc := NewMockIBucketCreator(ctrl)
	service := NewChapterService(repo, bc)

	uuid := "non-existent-uuid"

	repo.EXPECT().GetByUuid(gomock.Any(), uuid).Return(nil, fmt.Errorf("Такого uuid не существует!"))

	chapter, err := service.GetByUuid(context.Background(), uuid)
	if assert.Error(t, err) {
		assert.Nil(t, chapter)
		assert.Error(t, err)
	}
}

// TestChapterService_GetByUuid
func TestChapterService_List_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIChapterRepo(ctrl)
	bc := NewMockIBucketCreator(ctrl)
	service := NewChapterService(repo, bc)

	expectedChapters := dto.Chapters{
		{Uuid: "1", Number: 1, Title: "Chapter 1", PartUuid: "part-1", ReleaseUuid: nil},
		{Uuid: "2", Number: 2, Title: "Chapter 2", PartUuid: "part-2", ReleaseUuid: nil},
	}

	repo.EXPECT().List(gomock.Any()).Return(expectedChapters, nil)

	chapters, err := service.List(context.Background())
	if assert.NoError(t, err) {
		assert.Equal(t, expectedChapters, chapters)
	}
}

func TestChapterService_List_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIChapterRepo(ctrl)
	bc := NewMockIBucketCreator(ctrl)
	service := NewChapterService(repo, bc)

	repo.EXPECT().List(gomock.Any()).Return(nil, fmt.Errorf("Не удалось получить главы"))

	chapters, err := service.List(context.Background())
	if assert.Error(t, err) {
		assert.Nil(t, chapters)
		assert.Error(t, err)
	}
}

// TestChapterService_NextUuid
func TestChapterService_NextUuid_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIChapterRepo(ctrl)
	bc := NewMockIBucketCreator(ctrl)
	service := NewChapterService(repo, bc)

	uuid := "123e4567-e89b-12d3-a456-426655440000"
	expectedUuid := "456e7890-fedc-ba98-7654-321012345678"

	repo.EXPECT().NextUuid(gomock.Any(), uuid).Return(&expectedUuid, nil)

	nextUuid, err := service.NextUuid(context.Background(), uuid)
	if assert.NoError(t, err) {
		assert.Equal(t, &expectedUuid, nextUuid)
	}
}

func TestChapterService_NextUuid_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIChapterRepo(ctrl)
	bc := NewMockIBucketCreator(ctrl)
	service := NewChapterService(repo, bc)

	uuid := "non-existent-uuid"

	repo.EXPECT().NextUuid(gomock.Any(), uuid).Return(nil, fmt.Errorf("Нет такого uuid"))

	nextUuid, err := service.NextUuid(context.Background(), uuid)
	if assert.Error(t, err) {
		assert.Nil(t, nextUuid)
		assert.Error(t, err)
	}
}

func TestChapterService_NextUuid_NoNext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIChapterRepo(ctrl)
	bc := NewMockIBucketCreator(ctrl)
	service := NewChapterService(repo, bc)

	uuid := "123e4567-e89b-12d3-a456-426655440000"

	repo.EXPECT().NextUuid(gomock.Any(), uuid).Return(nil, fmt.Errorf("Нет следующего uuid"))

	nextUuid, err := service.NextUuid(context.Background(), uuid)
	if assert.Error(t, err) {
		assert.Nil(t, nextUuid)
		assert.Error(t, err)
	}
}

// TestChapterService_FirstUuid
func TestChapterService_FirstUuid_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIChapterRepo(ctrl)
	bc := NewMockIBucketCreator(ctrl)
	service := NewChapterService(repo, bc)

	expectedUuid := "123e4567-e89b-12d3-a456-426655440000"

	repo.EXPECT().FirstUuid(gomock.Any()).Return(expectedUuid, nil)

	uuid, err := service.FirstUuid(context.Background())
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if uuid != expectedUuid {
		t.Errorf("Expected UUID %s, got %s", expectedUuid, uuid)
	}
}

// TestChapterService_Create
func TestChapterService_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIChapterRepo(ctrl)
	bc := NewMockIBucketCreator(ctrl)
	service := NewChapterService(repo, bc)

	inputChapter := &dto.ChapterCreate{
		Title: "Chapter Title",
	}

	mockChapter := &dto.Chapter{
		Uuid:  "123e4567-e89b-12d3-a456-426655440000",
		Title: inputChapter.Title,
	}

	repo.EXPECT().Create(gomock.Any(), inputChapter).Return(mockChapter, nil)
	bc.EXPECT().CreateBucket(gomock.Any(), gomock.Any()).Return(nil)
	repo.EXPECT().Tx(gomock.Any()).Return(repo, &txMock{}, nil)

	createdChapter, err := service.Create(context.Background(), inputChapter)

	// Check the result
	if assert.NoError(t, err) {
		assert.Equal(t, mockChapter.Uuid, createdChapter.Uuid)
	}
}

// TestChapterService_Delete
func TestChapterService_Delete_Exist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIChapterRepo(ctrl)
	bc := NewMockIBucketCreator(ctrl)
	service := NewChapterService(repo, bc)

	uuid := "123e4567-e89b-12d3-a456-426655440000"

	repo.EXPECT().Delete(gomock.Any(), uuid).Return(nil)

	err := service.Delete(context.Background(), uuid)
	assert.NoError(t, err)
}

func TestChapterService_Delete_NotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockIChapterRepo(ctrl)
	bc := NewMockIBucketCreator(ctrl)
	service := NewChapterService(repo, bc)

	uuid := "123e4567-e89b-12d3-a456-426655440000"

	repo.EXPECT().Delete(gomock.Any(), uuid).Return(fmt.Errorf("Такого uuid не существует!"))

	err := service.Delete(context.Background(), uuid)
	assert.Error(t, err)
}
