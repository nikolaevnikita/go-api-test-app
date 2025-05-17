package server

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// MARK: Gzip Decompress Middleware

func GzipDecompressMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.Header.Get("Content-Encoding"), "gzip") {
			gz, err := gzip.NewReader(c.Request.Body)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			defer gz.Close()

			c.Request.Body = gz

			c.Request.ContentLength = -1
			c.Request.Header.Del("Content-Length")
			c.Request.Header.Del("Content-Encoding")
		}

		c.Next()
	}
}

// MARK: Gzip Compress Middleware

func GzipCompressMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.Contains(c.Request.Header.Get("Accept-Encoding"), "gzip") {
			c.Next()
			return
		}

		gz, err := gzip.NewWriterLevel(c.Writer, gzip.BestCompression)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer gz.Close()

		c.Writer.Header().Set("Content-Encoding", "gzip")

		c.Writer = &gzipResponseWriter{
			ResponseWriter: c.Writer,
			Writer:         gz,
		}

		c.Next()
	}
}

type gzipResponseWriter struct {
	gin.ResponseWriter
	io.Writer
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	contentType := w.Header().Get("Content-Type")
	contentTypeSupportsGzip := strings.Contains(contentType, "text/html") || strings.Contains(contentType, "application/json")
	if contentTypeSupportsGzip {
		return w.Writer.Write(b)
	}

	return w.ResponseWriter.Write(b)
}
