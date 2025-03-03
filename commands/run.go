package commands

import (
	"github.com/tothmate90/news-scraper/api"
	"github.com/tothmate90/news-scraper/config"
	"github.com/tothmate90/news-scraper/elasticsearch"
	"github.com/tothmate90/news-scraper/mysql"
)

// RunServer The main function setting up the server.
func RunServer(configFile string) error {
	config, err := config.ReadJSON(configFile)
	// MySQL section
	_, err = mysql.New(config.Conn)
	if err != nil {
		return err
	}
	// Elastic section
	elasticHandler, err := elasticsearch.New(config.Conn, "article")
	if err != nil {
		return err
	}
	// Api section
	apiHandler := api.New(elasticHandler, config)
	apiHandler.Listen()
	return err
}
