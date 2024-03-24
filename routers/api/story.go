package api

import (
	"net/http"
	"story-api/pkg/app"
	"story-api/pkg/e"
	"story-api/pkg/setting"
	"story-api/pkg/util"
	"story-api/service/story_service"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary Get all stories
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/stories [get]
func GetStories(c *gin.Context) {
	appG := app.Gin{C: c}

	storyService := story_service.Story{
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}
	stories, err := storyService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_STORIES_FAIL, nil)
		return
	}

	count, err := storyService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_STORY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": stories,
		"total": count,
	})
}

type AddStoryForm struct {
	Id         int    `form:"id" valid:"Required;Min(1)"`
	Name       string `form:"name" valid:"Required;MaxSize(100)"`
	StoryScore int    `form:"story_score" valid:"Required;Max(1000000000000000)"`
	// ConfidenceLevel represents the confidence level of a story.
	// It is an integer value that ranges from 0 to 5.
	// The higher the value, the higher the confidence in the story.
	ConfidenceLevel int `form:"confidence_level" valid:"Range(0,5)"`
	// Vermelha
	// Roxa
	// Laranja
	// Amarela
	// Azul
	// Verde
}

// @Summary Add story
// @Produce  json
// @Param id body int true "Id"
// @Param name body string true "Name"
// @Param story_score body int true "StoryScore"
// @Param confidence_level body int false "ConfidenceLevel"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/stories [post]
func AddStory(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddStoryForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	storyService := story_service.Story{
		Id:              form.Id,
		Name:            form.Name,
		StoryScore:      form.StoryScore,
		ConfidenceLevel: form.ConfidenceLevel,
	}
	exists, err := storyService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_STORY_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_STORY, nil)
		return
	}

	err = storyService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_STORY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type UpdateStoryForm struct {
	Id         int    `form:"id" valid:"Required;Min(1)"`
	Name       string `form:"name" valid:"Required;MaxSize(100)"`
	StoryScore int    `form:"story_score" valid:"Required;Max(1000000000000000)"`
	// ConfidenceLevel represents the confidence level of a story.
	// It is an integer value that ranges from 0 to 5.
	// The higher the value, the higher the confidence in the story.
	ConfidenceLevel int `form:"confidence_level" valid:"Range(0,6)"`
	// Vermelha => Jogo do vermelho (clicar até tudo ficar preto)
	// Roxa => Jogo do lápis (traduzir as palavras)
	// Laranja => Jogo do raio (traduzir as frases)
	// Amarela => Jogo da estrela (copiar as frases)
	// Dourado => Gincana do escrever (pintar os caracteres)
	// Azul => Jogo do sol (copiar o texto, traduzir todo o texto)
	// Verde
}

// @Summary Update story
// @Produce  json
// @Param id body int true "Id"
// @Param name body string true "Name"
// @Param story_score body int true "StoryScore"
// @Param confidence_level body int false "ConfidenceLevel"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/stories/{id} [put]
func UpdateStory(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = UpdateStoryForm{Id: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	storyService := story_service.Story{
		Id:              form.Id,
		Name:            form.Name,
		StoryScore:      form.StoryScore,
		ConfidenceLevel: form.ConfidenceLevel,
	}

	exists, err := storyService.ExistById()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_STORY_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_STORY, nil)
		return
	}

	err = storyService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_UPDATE_STORY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
