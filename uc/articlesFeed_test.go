package uc_test

import (
	"testing"

	"github.com/err0r500/go-realworld-clean/domain"
	mock "github.com/err0r500/go-realworld-clean/implem/mock.uc"
	"github.com/err0r500/go-realworld-clean/testData"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var art1 = domain.Article{Slug: "1"}
var art2 = domain.Article{Slug: "2"}
var art3 = domain.Article{Slug: "3"}
var art4 = domain.Article{Slug: "4"}

func TestInteractor_ArticlesFeed_happycases(t *testing.T) {
	rick := testData.User("rick")
	rick.FollowIDs = []string{testData.User("jane").Name}
	expectedArticles := []domain.Article{art1, art2, art3, art4}

	t.Run("most obvious", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		i := mock.NewMockedInteractor(mockCtrl)
		i.UserRW.EXPECT().GetByName(rick.Name).Return(&rick, nil)
		i.ArticleRW.EXPECT().GetByAuthorsNameOrderedByMostRecentAsc(rick.FollowIDs).Return(expectedArticles, nil).Times(1)

		articles, err := i.GetUCHandler().ArticlesFeed(rick.Name, 4, 0)
		assert.NoError(t, err)
		assert.Equal(t, expectedArticles, articles)

	})
	t.Run("offset& limit edgecases", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		i := mock.NewMockedInteractor(mockCtrl)
		i.UserRW.EXPECT().GetByName(rick.Name).Return(&rick, nil).AnyTimes()
		i.ArticleRW.EXPECT().GetByAuthorsNameOrderedByMostRecentAsc(rick.FollowIDs).Return(expectedArticles, nil).AnyTimes()

		t.Run("limit over length", func(t *testing.T) {
			articles, err := i.GetUCHandler().ArticlesFeed(rick.Name, 100, 0)
			assert.NoError(t, err)
			assert.Equal(t, expectedArticles, articles)
		})

		t.Run("offset over length", func(t *testing.T) {
			articles, err := i.GetUCHandler().ArticlesFeed(rick.Name, 4, 12)
			assert.NoError(t, err)
			assert.Equal(t, []domain.Article{}, articles)
		})

		t.Run("offset+limit over length", func(t *testing.T) {
			articles, err := i.GetUCHandler().ArticlesFeed(rick.Name, 3, 2)
			assert.NoError(t, err)
			assert.Equal(t, []domain.Article{art3, art4}, articles)
		})

		t.Run("negative limit", func(t *testing.T) {
			articles, err := i.GetUCHandler().ArticlesFeed(rick.Name, -100, 0)
			assert.NoError(t, err)
			assert.Equal(t, []domain.Article{}, articles)
		})

		t.Run("offset below 0", func(t *testing.T) {
			articles, err := i.GetUCHandler().ArticlesFeed(rick.Name, 4, -12)
			assert.NoError(t, err)
			assert.Equal(t, expectedArticles, articles)
		})

	})

}
