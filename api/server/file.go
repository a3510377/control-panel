// reference: https://gist.github.com/CJEnright/bc2d8b8dc0c1389a9feeddb110f822d7
package server

import (
	"compress/gzip"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

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

	s.RouterConfig.NoRouteHandlers = append(s.RouterConfig.NoRouteHandlers, func(c *gin.Context) {
		/* ---------- 404 page ---------- */
		UPath := c.Request.URL.Path
		if !strings.HasPrefix(UPath, "/") {
			UPath = "/" + UPath
		}
		UPath = path.Clean(UPath)

		f, err := fs.Open(UPath)
		if err != nil {
			if os.IsNotExist(err) {
				// to 404 page. suffix is `/` is important
				// if not, will be redirect to `${path}/404` ( is unlimited loop )
				c.Request.URL.Path = "/404/"
			}
		} else {
			f.Close()
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
