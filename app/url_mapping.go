package app

import (
	"repository.com/my_username/repo_name/controllers/ping"
	"repository.com/my_username/repo_name/controllers/shorturl_handlers"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/:short_path", shorturl_handlers.Redirect)
	router.GET("/:short_path/info", shorturl_handlers.Information)
	router.POST("/", shorturl_handlers.Create)

}
