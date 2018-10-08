package main

import (
	"github.com/nlnwa/maalfrid-aggregator/pkg/aggregator"
	"os"

	"github.com/inconshreveable/log15"
	"github.com/namsral/flag"
	"github.com/nlnwa/pkg/logfmt"
	"github.com/nlnwa/pkg/log"

	"github.com/nlnwa/maalfrid-aggregator/version"
	"github.com/nlnwa/maalfrid-language-detector/pkg/maalfrid"
)

func main() {
	port := 8672
	dbHost := "localhost"
	dbPort := 28015
	dbName := "test"
	dbUser := "admin"
	dbPassword := ""
	maalfridHost := "localhost"
	maalfridPort := 8672
	debug := false

	flag.IntVar(&port, "port", port, "gRPC server listening port")
	flag.StringVar(&dbHost, "db-host", dbHost, "database host")
	flag.IntVar(&dbPort, "db-port", dbPort, "database port")
	flag.StringVar(&dbName, "db-name", dbName, "database name")
	flag.StringVar(&dbUser, "db-user", dbUser, "database user")
	flag.StringVar(&dbName, "db-password", dbPassword, "database password")
	flag.StringVar(&maalfridHost, "maalfrid-host", maalfridHost, "maalfrid host")
	flag.IntVar(&maalfridPort, "maalfrid-port", maalfridPort, "maalfrid port")
	flag.BoolVar(&debug, "debug", debug, "enable debugging")
	flag.Parse()

	var logger log.Logger
	logger = log15.New()
	logHandler := log15.CallerFuncHandler(log15.StreamHandler(os.Stdout, logfmt.LogbackFormat()))
	// print stacktrace if debug option is true
	if debug {
		logger.(log15.Logger).SetHandler(log15.CallerStackHandler("%+v", logHandler))
	} else {
		logger.(log15.Logger).SetHandler(log15.LvlFilterHandler(log15.LvlInfo, logHandler))
	}

	logger.Info(version.String())

	store, err := aggregator.NewStore(aggregator.WithDatabase(dbHost, dbPort, dbName, dbUser, dbPassword))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	err = store.OpenStore()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer func() {
		if err = store.CloseStore(); err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}()
	maalfridClient := maalfrid.NewApiClient(maalfrid.WithAddress(maalfridHost, maalfridPort))
	api, err := aggregator.NewApi(aggregator.WithStore(store), aggregator.WithMaalfridClient(maalfridClient))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	err = aggregator.ServeApi(port, aggregator.WithLogger(&logger), aggregator.WithApi(api))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
