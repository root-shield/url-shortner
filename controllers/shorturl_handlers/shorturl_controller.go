package shorturl_handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"repository.com/my_username/repo_name/domain/shorturl"
	"repository.com/my_username/repo_name/services"
	"repository.com/my_username/repo_name/utils/errors"
)

func Create(c *gin.Context) {
	var short_url_req shorturl.ShortUrlRequest
	if err := c.ShouldBindJSON(&short_url_req); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	shorted_url, err := services.ShortUrlService.CreateShortUrl(short_url_req.Url)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, shorted_url)
}

func Information(c *gin.Context) {
	short_url_req := shorturl.ShortUrlRequest{ShortBase32: c.Param("short_path")}
	redirect_url, err := services.ShortUrlService.GetShortUrlByShortPath(short_url_req.ShortBase32)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, redirect_url)
}

func Redirect(c *gin.Context) {
	short_url_req := shorturl.ShortUrlRequest{ShortBase32: c.Param("short_path")}
	redirect_url, err := services.ShortUrlService.GetShortUrlByShortPath(short_url_req.ShortBase32)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	errIncremCount := services.ShortUrlService.IncrementShortUrlCount(short_url_req.ShortBase32)
	if errIncremCount != nil {
		c.JSON(err.Status, err)
		return
	}
	c.Redirect(http.StatusPermanentRedirect, redirect_url.Url)
}
