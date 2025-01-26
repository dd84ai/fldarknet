package web

import (
	"fmt"
	"html"
	"net/http"
	"strings"

	"github.com/darklab8/fl-darkstat/darkcore/builder"
	"github.com/darklab8/fl-darkstat/darkcore/core_types"
	"github.com/darklab8/fl-darkstat/darkcore/settings/logus"
	"github.com/darklab8/fl-darkstat/darkcore/web/registry"
	"github.com/darklab8/go-typelog/typelog"
	"github.com/darklab8/go-utils/utils/utils_types"
)

const UrlStatic core_types.Url = "/"

func NewEndpointStatic(w *Web) *registry.Endpoint {
	return &registry.Endpoint{
		Url: UrlStatic,
		Handler: func(resp http.ResponseWriter, req *http.Request) {
			if w.AppDataMutex != nil {
				w.AppDataMutex.Lock()
				defer w.AppDataMutex.Unlock()
			}
			switch req.Method {
			case http.MethodOptions:
			case http.MethodGet:

				requested := req.URL.Path[1:]

				requested = strings.ReplaceAll(requested, "/", core_types.PATH_SEPARATOR)

				replaced_root := strings.Replace(requested, w.site_root[1:], "", 1)
				logus.Log.Debug("replacing static root",
					typelog.String("requested", requested),
					typelog.String("replaced_root", replaced_root),
					typelog.String("w.site_root[1:]", w.site_root[1:]),
				)
				requested = replaced_root

				if requested == "" {
					requested = "index.html"
				}

				logger := logus.Log.WithFields(
					typelog.String("requested_path", requested),
					typelog.Int("files_count", len(w.filesystems[0].Files)),
				)

				logger.Info("having get request")

				var content builder.MemFile
				var ok bool
				for _, filesystem := range w.filesystems {

					content, ok = filesystem.Files[utils_types.FilePath(requested)]
					if ok {
						break
					}
				}

				if strings.Contains(requested, ".css") {
					resp.Header().Set("Content-Type", "text/css; charset=utf-8")
				} else if strings.Contains(requested, ".html") {
					resp.Header().Set("Content-Type", "text/html; charset=utf-8")
				} else if strings.Contains(requested, ".js") {
					resp.Header().Set("Content-Type", "application/javascript; charset=utf-8")
				}

				if ok {
					fmt.Fprint(resp, string(content.Render()))
				} else {
					resp.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(resp, "content is not found at %s!, %q", req.URL, html.EscapeString(requested))
					logus.Log.Error("content is not found")
				}

			default:
				http.Error(resp, "Method not allowed", http.StatusMethodNotAllowed)
			}
		},
	}
}
