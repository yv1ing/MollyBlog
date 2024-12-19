package mApp

func (ma *MApp) loadRoutes() {
	ma.engine.Use(ma.AuthMiddleware)

	ma.engine.GET("/", ma.IndexHandler)
	ma.engine.GET("/rss", ma.RSSHandler)
	ma.engine.GET("/about", ma.AboutHandler)
	ma.engine.GET("/search", ma.SearchHandler)
	ma.engine.GET("/archive", ma.ArchiveHandler)
	ma.engine.GET("/post/:hash", ma.PostHandler)
	ma.engine.GET("/tag/:hash", ma.TagHandler)
	ma.engine.GET("/category/:hash", ma.CategoryHandler)

	// update's endpoint
	ma.engine.POST(ma.Config.UpdateEndpoint, ma.UpdateBlogHandler)
}
