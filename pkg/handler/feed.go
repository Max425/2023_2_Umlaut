package handler

import (
	"net/http"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/utils"
	"github.com/golang/protobuf/ptypes"
)

// @Summary get user for feed
// @Tags feed
// @ID feed
// @Accept  json
// @Produce  json
// @Param min_age query integer false "Minimum age filter"
// @Param max_age query integer false "Maximum age filter"
// @Param tags query string false "Tags filter"
// @Success 200 {object} ClientResponseDto[model.User]
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/feed [get]
func (h *Handler) feed(w http.ResponseWriter, r *http.Request) {
	user, err := h.feedMicroservice.Feed(
		r.Context(),
		utils.ParseQueryParams(r),
	)

	if err != nil {
		statusCode, message := utils.ParseError(err)
		newErrorClientResponseDto(r.Context(), w, statusCode, message)
		return
	}
	birthday, err := ptypes.Timestamp(user.Birthday)
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}
	preferGender := int(user.PreferGender)
	age := int(user.Age)

	NewSuccessClientResponseDto(r.Context(), w, model.User{
		Id:           int(user.Id),
		Name:         user.Name,
		PreferGender: &preferGender,
		Description:  &user.Description,
		Age:          &age,
		Looking:      &user.Looking,
		Education:    &user.Education,
		Hobbies:      &user.Hobbies,
		Birthday:     &birthday,
		Tags:         &user.Tags,
		ImagePaths:   &user.ImagePaths,
	})
}
