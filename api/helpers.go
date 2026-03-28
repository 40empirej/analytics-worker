package analyticsworker

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Headers struct {
	Host     string
	Referer  string
	UserAgent string
	Accept    string
}

type RequestOptions struct {
	Method  string
	Headers Headers
	Body    string
	Timeout time.Duration
}

type Response struct {
	Status int
	Body   string
}

func getHeaders(host string) Headers {
	return Headers{
		Host:     host,
		Referer:  host,
		UserAgent: "Analytics-Worker/1.0",
		Accept:   "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	}
}

func getResponse(r *http.Request) Response {
	resp, err := http.DefaultClient.Do(r)
	if err!= nil {
		log.Println(err)
		return Response{Status: http.StatusInternalServerError}
	}
	defer resp.Body.Close()

	status := resp.StatusCode
	body, err := strconv.unquote(resp.Header.Get("Content-Encoding"))
	if err!= nil {
		body = resp.Body.String()
	}
	return Response{Status: status, Body: body}
}

func extractDomain(host string) string {
	parts := strings.Split(host, ".")
	return strings.Join(parts[len(parts)-2:], ".")
}