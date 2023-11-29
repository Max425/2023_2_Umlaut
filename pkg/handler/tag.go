package handler

import "net/http"

// @Summary get all tags
// @Tags tag
// @ID tag
// @Accept  json
// @Produce  json
// @Success 200 {object} ClientResponseDto[[]string]
// @Failure 401,500 {object} ClientResponseDto[string]
// @Router /api/v1/tag [get]
func (h *Handler) getAllTags(w http.ResponseWriter, r *http.Request) {
	tags, err := h.services.Tag.GetAllTags(r.Context())
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessClientResponseArrayDto(r.Context(), w, tags)
}
