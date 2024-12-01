package service

import (
	"fmt"
	"time"
)

func (articleService *ArticleServiceImpl) Update(description, image string, id int) error {
	ifExist, err := articleService.ArticleRepository.GetById(id)
	if err != nil {
		return fmt.Errorf("failed to check is article already exist:%s", err.Error())
	}

	if ifExist == nil {
		return fmt.Errorf("article not found")
	}

	//update updated at
	wib, err := time.LoadLocation("Asia/Jakarta") // WIB (UTC+7)
	if err != nil {
		return fmt.Errorf("error loading WIB location: %s", err.Error())
	}

	now := time.Now().In(wib)
	updatedAt := now

	err = articleService.ArticleRepository.Update(description, image, id, updatedAt)
	if err != nil {
		return fmt.Errorf("error when updating : %s", err.Error())
	}

	return nil
}
