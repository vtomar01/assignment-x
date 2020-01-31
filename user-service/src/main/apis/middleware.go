package apis

type MiddleWare func(handler RequestHandler) RequestHandler
