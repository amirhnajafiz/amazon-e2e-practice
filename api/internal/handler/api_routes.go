package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// getUrls handles the request to retrieve all URLs from the database.
func (h Handler) getUrls(c echo.Context) error {
	urls, err := h.DB.GetUrls()
	if err != nil {
		log.Printf("failed to retrieve URLs: %v\n", err)
		return c.String(http.StatusInternalServerError, "failed to retrieve URLs")
	}

	return c.JSON(http.StatusOK, urls)
}

// insertSession handles the request to insert a new session into the database.
func (h Handler) insertSession(c echo.Context) error {
	var req struct {
		UrlID int64 `json:"url_id"`
	}
	if err := c.Bind(&req); err != nil {
		log.Printf("failed to bind request: %v\n", err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	if err := h.DB.InsertSession(req.UrlID); err != nil {
		log.Printf("failed to insert session: %v\n", err)
		return c.String(http.StatusInternalServerError, "failed to insert session")
	}

	return c.NoContent(http.StatusCreated)
}

// getTopUrls handles the request to retrieve the top URLs by session count.
func (h Handler) getTopUrls(c echo.Context) error {
	limit := 10
	if l := c.QueryParam("limit"); l != "" {
		if _, err := fmt.Sscanf(l, "%d", &limit); err != nil {
			log.Printf("invalid limit parameter: %v\n", err)
			return c.String(http.StatusBadRequest, "invalid limit parameter")
		}
	}

	urls, err := h.DB.GetTopUrlsBySessionCount(limit)
	if err != nil {
		log.Printf("failed to retrieve top URLs: %v\n", err)
		return c.String(http.StatusInternalServerError, "failed to retrieve top URLs")
	}

	return c.JSON(http.StatusOK, urls)
}

// getSessionsCount handles the request to retrieve the count of sessions for a specific URL ID.
func (h Handler) getSessionsCount(c echo.Context) error {
	urlID := c.Param("url_id")
	if urlID == "" {
		return c.String(http.StatusBadRequest, "url_id is required")
	}

	var urlIDInt int64
	if _, err := fmt.Sscanf(urlID, "%d", &urlIDInt); err != nil {
		log.Printf("invalid url_id parameter: %v\n", err)
		return c.String(http.StatusBadRequest, "invalid url_id parameter")
	}

	count, err := h.DB.GetSessionsCountByUrlID(urlIDInt)
	if err != nil {
		log.Printf("failed to get sessions count for URL ID %d: %v\n", urlIDInt, err)
		return c.String(http.StatusInternalServerError, "failed to get sessions count")
	}

	return c.JSON(http.StatusOK, map[string]int{"count": count})
}
