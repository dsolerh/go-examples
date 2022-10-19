package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/dsolerh/examples/concurrency.course/finalProject/pkg/data"
)

var testApp Config

func TestMain(m *testing.M) {
	gob.Register(data.User{})

	// setup session
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	// create loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// set up application config
	testApp = Config{
		Session:       session,
		DB:            nil,
		InfoLog:       infoLog,
		ErrorLog:      errorLog,
		Wait:          &sync.WaitGroup{},
		ErrorChan:     make(chan error),
		ErrorDoneChan: make(chan bool),
		// Models:        data.New(db),
	}

	// create a dummy mailer

	testApp.Mailer = MailServer{
		Domain:      "localhost",
		Host:        "localhost",
		Port:        1025,
		Encryption:  "none",
		FromName:    "info",
		FromAddress: "info@myconpany.com",
		Wait:        testApp.Wait,
		ErrorChan:   make(chan error),
		MailerChan:  make(chan Message, 100),
		DoneChan:    make(chan bool),
	}

	go func() {
		select {
		case <-testApp.Mailer.MailerChan:
		case <-testApp.Mailer.ErrorChan:
		case <-testApp.Mailer.DoneChan:
			return
		}

	}()

	go func() {
		for {
			select {
			case err := <-testApp.ErrorChan:
				testApp.ErrorLog.Println(err)
			case <-testApp.ErrorDoneChan:
				return
			}
		}
	}()

	os.Exit(m.Run())
}
