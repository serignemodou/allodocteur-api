package initialize

import (
	env "allodocteur-api/pkg/environment"
	"log"
	"net/http"
	"strconv"
)

func init() {
	initDB()
	log.Println("listen on port", env.ServerURL)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(env.ServerPort), nil))
}
