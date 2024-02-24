package main

import (
	"github.com/adiatma85/gg-project/src/business/domain"
	"github.com/adiatma85/gg-project/src/business/usecase"
	"github.com/adiatma85/gg-project/src/handler"
	"github.com/adiatma85/gg-project/utils/config"
	"github.com/adiatma85/own-go-sdk/configreader"
	"github.com/adiatma85/own-go-sdk/instrument"
	"github.com/adiatma85/own-go-sdk/jwtAuth"
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/parser"
	"github.com/adiatma85/own-go-sdk/sql"
)

// @contact.name   Rahmadhani Lucky Adiatma

// @securitydefinitions.apikey BearerAuth
// @in header
// @name Authorization

const (
	configfile   string = "./etc/cfg/conf.json"
	templatefile string = "./etc/tpl/conf.template.json"
)

func main() {
	// Build the config
	// Assume the config is exist in the first place

	// Read the Config first
	cfg := config.Init()
	configreader := configreader.Init(configreader.Options{
		ConfigFile: configfile,
	})
	configreader.ReadConfig(&cfg)

	// init logger
	log := log.Init(cfg.Log)

	// init the instrument
	instr := instrument.Init(cfg.Instrument)

	// Init the DB
	db := sql.Init(cfg.SQL, log, instr)

	// init the parser
	parsers := parser.InitParser(log, cfg.Parser)

	// Init the jwt
	jwt := jwtAuth.Init(cfg.JwtAuth)

	// Init the domain
	d := domain.Init(domain.InitParam{Log: log, Db: db, Json: parsers.JSONParser()})

	// Init the usecase
	uc := usecase.Init(usecase.InitParam{Log: log, Dom: d, JwtAuth: jwt})

	// Init the GIN
	rest := handler.Init(handler.InitParam{Conf: cfg.Gin, Json: parsers.JSONParser(), Log: log, Uc: uc, Instrument: instr, JwtAuth: jwt})

	rest.Run()
}
