package swagger

import (
	stlog "log"
	"log/slog"
	"net/http"
	"path/filepath"
	"sync"

	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	auth         = "pkg/pb/api/auth"
	personalinfo = "pkg/pb/api/personalinfo"
)

func InitSwagger(wg *sync.WaitGroup, log *slog.Logger) {
	defer wg.Done()

	router := http.NewServeMux()


	swaggerDirs := map[string]string{
		"authService":         auth,
		"personalInfoService": personalinfo,
	}


	for key, dir := range swaggerDirs {
		swaggerPath := filepath.Join(dir, "swagger.json")


		router.HandleFunc("/swagger/"+key+"/swagger.json", func(path string) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				http.ServeFile(w, r, path)
			}
		}(swaggerPath))


		router.Handle("/swagger/"+key+"/", httpSwagger.Handler(
			httpSwagger.URL("/swagger/"+key+"/swagger.json"),
		))
	}

	stlog.Fatal(http.ListenAndServe(":8082", router))
}
