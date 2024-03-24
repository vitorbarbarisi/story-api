package story_service

import (
	"story-api/models"
)

type Story struct {
	Id         int
	Name       string
	StoryScore int
	// ConfidenceLevel represents the confidence level of a story.
	// It is an integer value that ranges from 0 to 5.
	// The higher the value, the higher the confidence in the story.
	ConfidenceLevel int

	PageNum  int
	PageSize int
}

func (s *Story) ExistByName() (bool, error) {
	return models.ExistStoryByName(s.Name)
}

func (s *Story) ExistById() (bool, error) {
	return models.ExistStoryByID(s.Id)
}

func (s *Story) Add() error {
	return models.AddStory(s.Id, s.Name, s.StoryScore, s.ConfidenceLevel)
}

func (s *Story) Edit() error {
	data := make(map[string]interface{})
	data["confidence_level"] = s.ConfidenceLevel

	return models.EditStory(s.Id, data)
}

func (s *Story) Count() (int, error) {
	return models.GetStoryTotal(s.getMaps())
}

func (s *Story) GetAll() ([]models.Story, error) {
	tags, err := models.GetStories(s.PageNum, s.PageSize, s.getMaps())
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (s *Story) getMaps() map[string]interface{} {
	return make(map[string]interface{})
}
