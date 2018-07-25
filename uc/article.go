package uc

import "github.com/err0r500/go-realworld-clean/domain"

func (i interactor) ArticlePost(article domain.Article) (*domain.Article, error) { return nil, nil }
func (i interactor) ArticlePut(slug string, article domain.Article) (*domain.Article, error) {
	return nil, nil
}
func (i interactor) ArticleGet(slug string) (*domain.Article, error) { return nil, nil }
func (i interactor) ArticleDelete(slug string) error                 { return nil }
