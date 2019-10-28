package main

import (
	"flag"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/maratgaliev/golangjobs/api"
	"github.com/maratgaliev/golangjobs/config"
	"github.com/maratgaliev/golangjobs/service"
	"github.com/maratgaliev/golangjobs/store"
)

func main() {
	var err error
	log.SetFormatter(&log.JSONFormatter{})
	configFile := flag.String("c", "config.yaml", "Path to config file")
	flag.Parse()
	var conf *config.Config
	if conf, err = config.NewConfig(*configFile); err != nil {
		log.Fatal(err)
	}
	log.SetLevel(log.Level(conf.LogLevel))
	log.Info("Starting service with configuration: ", conf.ConfigFile)
	store, err := store.NewStore(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()
	log.Info("Store created successfully")
	js := service.NewJobService(store)
	log.Info("Services created successfully")
	api := api.ApiRunner(conf, js)
	log.WithField("address", api.GetApiInfo().Address).
		WithField("mw", api.GetApiInfo().MW).
		WithField("routes", api.GetApiInfo().Routes).
		Info("Starting api")
	log.Fatal(api.Start())
}
