package service

import "github.com/simiancreative/services/server/req"

type Handler func(req.Context) (interface{}, error)

type Collection []Config

type Config struct {
	Method  string
	Path    string
	Handler Handler
	Input   func() interface{}
	Auth    func(req.Context) error
	Before  func(req.Context) error
	After   func(req.Context)
}
