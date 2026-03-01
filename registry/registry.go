package registry

import "github.com/Meduzz/sture/model"

var (
	storage = make(map[model.Kind]model.Engine)
)

func Register(kind model.Kind, engine model.Engine) {
	storage[kind] = engine
}

func EngineFor(kind model.Kind) (model.Engine, bool) {
	engine, ok := storage[kind]

	return engine, ok
}
