package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Story struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	StoryScore int    `json:"story_score"`
	// ConfidenceLevel represents the confidence level of a story.
	// It is an integer value that ranges from 0 to 5.
	// The higher the value, the higher the confidence in the story.
	ConfidenceLevel int `json:"confidence_level"`
}

// ExistStoryByName checks if there is a story with the same name
func ExistStoryByName(name string) (bool, error) {
	var story Story
	fmt.Printf("db!\n")
	if db == nil {
		fmt.Println("db is nil")
	} else {
		fmt.Println("db is not nil")
	}
	err := db.Select("id").Where("name = ?", name).First(&story).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if story.Id > 0 {
		return true, nil
	}

	return false, nil
}

// AddStory Add a Story
func AddStory(id int, name string, storyScore int, confidenceLevel int) error {
	story := Story{
		Id:              id,
		Name:            name,
		StoryScore:      storyScore,
		ConfidenceLevel: confidenceLevel,
	}
	if err := db.Create(&story).Error; err != nil {
		return err
	}

	return nil
}

// GetStories gets a list of stories based on paging and constraints
func GetStories(pageNum int, pageSize int, maps interface{}) ([]Story, error) {
	var (
		stories []Story
		err     error
	)

	if pageSize > 0 && pageNum >= 0 {
		err = db.Where(maps).Order("story_score desc").Offset(pageNum).Limit(pageSize).Find(&stories).Error
	} else {
		err = db.Where(maps).Order("story_score desc").Find(&stories).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return stories, nil
}

// GetStoryTotal counts the total number of stories based on the constraint
func GetStoryTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Story{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// ExistStoryByID determines whether a Story exists based on the ID
func ExistStoryByID(id int) (bool, error) {
	var story Story
	err := db.Select("id").Where("id = ?", id).First(&story).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if story.Id > 0 {
		return true, nil
	}

	return false, nil
}

// EditStory modifies a single story
func EditStory(id int, data interface{}) error {
	if err := db.Model(&Story{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
