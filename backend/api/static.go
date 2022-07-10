package api

import (
	"cafe/config"
	"github.com/gin-contrib/static"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

func (a *Api) handleStaticFiles() {
	a.Router.LoadHTMLFiles(config.TemplatesDir + "index.html")
	a.serveFoldersInTemplates()
}

func (a *Api) serveFoldersInTemplates() {
	_ = filepath.WalkDir(config.TemplatesDir, func(path string, info os.DirEntry, err error) error {
		if info.IsDir() && info.Name() != strings.TrimSuffix(config.TemplatesDir, "/") {
			a.Router.Use(static.Serve("/"+info.Name(), static.LocalFile(config.TemplatesDir+info.Name(), false)))
			logrus.WithField("folder", info.Name()).Debug("Serve static folder")
		}
		return err
	})
}
