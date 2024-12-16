package mApp

func (ma *MApp) loadRoutes() {
	ma.engine.GET("/", ma.IndexHandler)
	ma.engine.GET("/archive", ma.ArchiveHandler)
	ma.engine.GET("/post/:hash", ma.PostHandler)
	ma.engine.GET("/tag/:hash", ma.TagHandler)

	ma.engine.PUT("/update", ma.UpdateBlogHandler)
}
