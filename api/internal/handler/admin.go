package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// createUrl handles the request to create a new shortened URL.
func (h Handler) createUrl(c echo.Context) error {
	var req struct {
		ShortenedID string `json:"shortened_id"`
		Address     string `json:"address"`
		Description string `json:"description"`
	}
	if err := c.Bind(&req); err != nil {
		return c.String(400, "invalid request")
	}

	// validate required fields
	if req.ShortenedID == "" || req.Address == "" {
		return c.String(400, "shortened_id and address are required")
	}

	if err := h.DB.InsertUrl(req.ShortenedID, req.Address, req.Description); err != nil {
		return c.String(500, "failed to insert URL")
	}

	return c.NoContent(201)
}

// deleteUrl handles the request to delete a shortened URL by its ID.
func (h Handler) deleteUrl(c echo.Context) error {
	urlID := c.Param("url_id")
	if urlID == "" {
		return c.String(400, "url_id is required")
	}

	// convert urlID to int64
	var id int64
	if _, err := fmt.Sscanf(urlID, "%d", &id); err != nil {
		return c.String(400, "invalid url_id format")
	}

	if err := h.DB.DeleteUrl(id); err != nil {
		return c.String(500, "failed to delete URL")
	}
	if err := h.DB.DeleteSessionByUrlID(id); err != nil {
		return c.String(500, "failed to delete sessions for URL")
	}

	return c.NoContent(204)
}
