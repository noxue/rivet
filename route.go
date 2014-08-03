package rivet

// base route
type base struct {
	rivet    Riveter
	handlers []Handler
}

func (b *base) Rivet(rivet Riveter) {
	b.rivet = rivet
}

func (b *base) Handlers(handlers ...Handler) {
	b.handlers = handlers
}

func (b *base) Apply(params Params, context Context) {
	if b == nil {
		_, req := context.Source()
		panic(req.URL.Path)
	}
	if context != nil {
		if b.rivet == nil {
			context.Invoke(params, b.handlers...)
		} else {
			b.rivet.Context(context.Source()).Invoke(params, b.handlers...)
		}
	}
}