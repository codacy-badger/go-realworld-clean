package uc

import (
	"github.com/err0r500/go-realworld-clean/domain"
)

type Filters struct {
	AuthorFilter    *string
	TagFilter       *string
	FavoritedFilter *bool
}

func (i interactor) GetArticles(limit, offset int, filters Filters) (domain.ArticleCollection, int, error) {
	if limit <= 0 {
		return domain.ArticleCollection{}, 0, nil
	}

	articles, err := i.articleRW.GetRecentFiltered(filters)
	if err != nil {
		return nil, 0, err
	}

	return domain.ArticleCollection(articles).ApplyLimitAndOffset(limit, offset), len(articles), nil
}
