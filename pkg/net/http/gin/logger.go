package gin

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const (
	FieldKeyRequest  = "request"
	FieldKeyResponse = "response"

	Message = "RESPONSE_SENT"
)

type Request struct {
	Method      string          `json:"method"`
	URI         string          `json:"uri"`
	Body        json.RawMessage `json:"body"`
	RemoteAddr  string          `json:"remote_addr"`
	UserAgent   string          `json:"user_agent"`
	RequestUUID string          `json:"request_uuid"`
}

type Response struct {
	Status int             `json:"status"`
	Body   json.RawMessage `json:"body"`
}

type bodyWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

type LoggerMiddlewareConfig struct {
	uUIDHeaderKey string
	ignoredPaths  map[string]bool
}

// NewLoggerMiddlewareConfig instantiate gin logger middleware
func NewLoggerMiddlewareConfig(
	uUIDHeaderKey string,
	ignoredPaths []string,
) *LoggerMiddlewareConfig {
	l := &LoggerMiddlewareConfig{
		uUIDHeaderKey: uUIDHeaderKey,
		ignoredPaths:  map[string]bool{},
	}

	for _, p := range ignoredPaths {
		l.ignoredPaths[p] = true
	}

	return l
}

func LoggerMiddleware(config *LoggerMiddlewareConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBodyBytes []byte
		var requestBodyBytesLogging *bytes.Buffer
		var responseBodyWriter bodyWriter

		// for ignored paths don't log
		if _, ok := config.ignoredPaths[c.FullPath()]; ok {
			c.Next()
			return
		}

		// if req body not null, read req body bytes and
		// create a requestBodyBytesLogging from read req body bytes
		if c.Request.Body != nil {
			requestBodyBytes, _ = io.ReadAll(c.Request.Body)
			requestBodyBytesLogging = bytes.NewBuffer(requestBodyBytes)
		}

		// providing Close() from io.NopCloser to *gin.Context.Request.Body
		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBodyBytes))

		// create new response body writer
		responseBodyWriter = bodyWriter{
			bodyBuf:        bytes.NewBufferString(""),
			ResponseWriter: c.Writer}

		// assign response body writer to the context gin.ResponseWriter
		c.Writer = responseBodyWriter

		// call next to wrap the request and handle context after handling request if finished
		c.Next()

		var body []byte

		// get buffered bytes from request body writer to body
		if requestBodyBytesLogging != nil && requestBodyBytesLogging.Len() > 0 {
			body = requestBodyBytesLogging.Bytes()
		} else {
			body = []byte("{}")
		}

		// create request to log it
		request := Request{
			Method:      c.Request.Method,
			URI:         c.Request.RequestURI,
			Body:        json.RawMessage(body),
			RemoteAddr:  c.Request.RemoteAddr,
			UserAgent:   c.Request.UserAgent(),
			RequestUUID: c.Request.Header.Get(config.uUIDHeaderKey),
		}

		// create json from request struct
		req, err := json.Marshal(request)

		// log potential errors
		if err != nil {
			log.Error().Err(err)
		}

		// create a response struct for logging
		response := &Response{
			Status: c.Writer.Status(),
			Body: func() json.RawMessage {
				// trying to marshall empty body - empty byte slice results in err
				if len(responseBodyWriter.bodyBuf.Bytes()) < 1 {
					return nil
				}

				return responseBodyWriter.bodyBuf.Bytes()
			}(),
		}

		// parse json from response
		res, err := json.Marshal(response)

		// log error if parsing fails
		if err != nil {
			log.Error().Err(err)
		}

		// log request, response, and response sent message
		log.Info().
			Str(FieldKeyResponse, string(req)).
			Str(FieldKeyRequest, string(res)).
			Msg(Message)
	}
}
