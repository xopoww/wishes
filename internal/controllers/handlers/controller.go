package handlers

import "github.com/xopoww/wishes/internal/service"

type ApiController struct {
	t Trace
	s service.Service
}

func NewApiController(t Trace, s service.Service) *ApiController {
	return &ApiController{
		t: t,
		s: s,
	}
}
