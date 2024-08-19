package swagger

import (
	stlog "log"
	"log/slog"
	"net/http"
	"sync"

	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	auth = "pkg/pb/api/auth"
	home = "pkg/pb/api/home"
)

func InitSwagger(wg *sync.WaitGroup, log *slog.Logger) {
	defer wg.Done()

	router := http.NewServeMux()

	swaggerDirs := map[string]string{
		"auth": auth,
		"home": home,
	}

	for key, dir := range swaggerDirs {

		log.Info(key)
		log.Info(dir)
		router.HandleFunc("swagger/"+key+"/swagger.json", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "docs/swagger.json")
		})

		router.Handle("/swagger/"+key+"/", httpSwagger.Handler(
			httpSwagger.URL("/swagger/"+key+"/swagger.json"),
		))
	}

	stlog.Fatal(http.ListenAndServe(":8082", router))
}
