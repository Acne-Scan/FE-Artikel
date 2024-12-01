package repository

func (articleRepository *ArticleRepositoryImpl) Delete(id int) error {
	_, err := articleRepository.DB.Exec("delete from article where article_id=?", id)
	if err != nil {
		return err
	}

	return nil
}
