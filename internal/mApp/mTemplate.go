package mApp

import "fmt"

func (ma *MApp) loadTemplates() {
	assetsPath := fmt.Sprintf("%s/assets", ma.Config.Template)
	htmlPath := fmt.Sprintf("%s/html/*.html", ma.Config.Template)

	ma.engine.Static("/assets", assetsPath)
	ma.engine.LoadHTMLGlob(htmlPath)
}
