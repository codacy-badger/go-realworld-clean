package uc

import (
	"github.com/err0r500/go-realworld-clean/domain"
)

func (i interactor) ArticlesFeed(username string, limit, offset int) ([]domain.Article, error) {
	if limit < 0 {
		return []domain.Article{}, nil
	}

	user, err := i.userRW.GetByName(username)
	if err != nil {
		return nil, err
	}
	articles, err := i.articleRW.GetByAuthorsNameOrderedByMostRecentAsc(user.FollowIDs)
	if err != nil {
		return nil, err
	}

	articlesSize := len(articles)
	min := offset
	if min < 0 {
		min = 0
	}

	if min > articlesSize {
		return []domain.Article{}, nil
	}

	max := min + limit
	if max > articlesSize {
		max = articlesSize
	}

	return articles[min:max], nil
}
