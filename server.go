package gocouchlib

import (
	"net/url"
)

type Server struct {
	Url      string
	UserInfo *url.Userinfo
}

func (s *Server) FullUrl() string {
	fullUrl, _ := url.Parse(s.Url)
	if s.UserInfo != nil {
		fullUrl.User = s.UserInfo
	}

	return fullUrl.String()
}

func (s *Server) endpoint(api string) string {
	return s.FullUrl() + api
}

func (s *Server) Info() JsonObj {
	couchResp, _ := httpClient.Get(s.endpoint("/"), nil)
	return couchResp.Json
}

func (s *Server) AllDbs() JsonObj {
	couchResp, _ := httpClient.Get(s.endpoint("/_all_dbs"), nil)
	return couchResp.Json
}