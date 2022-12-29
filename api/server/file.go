// reference: https://gist.github.com/CJEnright/bc2d8b8dc0c1389a9feeddb110f822d7
package server

import (
	"compress/gzip"
	"io"
	"io/fs"
	"net/http"
	pathLib "path"
	"regexp"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var checkDynamicRoute = regexp.MustCompile(`/\[[^/]*\]`)

var gzPool = sync.Pool{
	New: func() any {
		w := gzip.NewWriter(io.Discard)
		gzip.NewWriterLevel(w, gzip.BestCompression)
		return w
	},
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *gzipResponseWriter) WriteHeader(status int) {
	w.ResponseWriter.Header().Del("Content-Length")
	w.ResponseWriter.WriteHeader(status)
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (s *Server) AddFileHandler(dir fs.FS) {
	fs := http.FS(dir)
	fileServer := http.FileServer(fs)

	routes := getRoutes(dir)
	s.RouterConfig.NoRouteHandlers = append(s.RouterConfig.NoRouteHandlers, func(c *gin.Context) {
		/* ---------- 404 page ---------- */
		UPath := pathLib.Clean(c.Request.URL.Path)

		if ok, path := routes.HasIs(UPath); ok {
			// suffix is `/` is important
			// if not, will be redirect to `${path}/${path}` ( is unlimited loop )
			c.Request.URL.Path = path + "/"
		} else {
			// to 404 page. suffix is `/` is important
			// if not, will be redirect to `${path}/404` ( is unlimited loop )
			c.Request.URL.Path = "/404/"
		}

		/* ---------- gzip ---------- */
		if !strings.Contains(c.GetHeader("Accept-Encoding"), "gzip") {
			fileServer.ServeHTTP(c.Writer, c.Request)
			return
		}

		// support gzip
		c.Header("Content-Encoding", "gzip")

		gz := gzPool.Get().(*gzip.Writer)
		defer gzPool.Put(gz)

		gz.Reset(c.Writer)
		defer gz.Close()

		fileServer.ServeHTTP(&gzipResponseWriter{ResponseWriter: c.Writer, Writer: gz}, c.Request)
	})
}

type route map[string]*route

// get routes from embed files
// check path is match `/\[[^/]*\]` ( for next.js export path format )
func getRoutes(dir fs.FS) route {
	dPaths := route{}

	fs.WalkDir(dir, ".", func(path string, file fs.DirEntry, _ error) (err error) {
		if file.IsDir() {
			if strings.HasPrefix(path, "/") { // Make sure the regex test is correct
				path = "/" + path
			}
			if checkDynamicRoute.MatchString(path) {
				var t *route
				for i, p := range strings.Split(path, "/") {
					if i == 0 {
						route := &route{}
						dPaths[p] = route
						t = route
					} else {
						r := &route{}
						(*t)[p] = r
						t = r
					}
				}
			}
		}

		return // return nil
	})

	return dPaths
}

// TODO add [...xxx] support
// check path is match route
func (s route) HasIs(path string) (bool, string) {
	var t route

	resultPath := ""
	paths := strings.Split(strings.TrimSuffix(strings.TrimPrefix(path, "/"), "/"), "/")
	for i, p := range paths {
		if i == 0 {
			if _, ok := s[p]; ok {
				t = *s[p]             // default route map
				resultPath += "/" + p // add current path
				continue
			} else {
				return false, ""
			}
		}

		if _, ok := t[p]; ok {
			resultPath += "/" + p // add current path
			t = *t[p]

			if i == len(paths)-1 {
				return true, resultPath
			}
			continue
		}

		check := false
		// check dynamic route
		for key := range t {
			if strings.HasPrefix(key, "[") && strings.HasSuffix(key, "]") {
				t = *t[key]             // next route map
				resultPath += "/" + key // add current path
				if i == len(paths)-1 {
					return true, resultPath
				}
				check = true
				break
			}
		}
		if !check {
			return false, ""
		}
	}

	if resultPath != "" {
		return true, resultPath
	} else {
		return false, ""
	}
}
