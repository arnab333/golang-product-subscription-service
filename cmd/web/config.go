package main

import (
	"database/sql"
	"log"
	"product_subscription_service/data"
	"sync"

	"github.com/alexedwards/scs/v2"
)

type Config struct {
	Session       *scs.SessionManager
	DB            *sql.DB
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	WG            *sync.WaitGroup
	Models        data.Models
	Mailer        Mail
	ErrorChan     chan error
	ErrorChanDone chan bool
}
