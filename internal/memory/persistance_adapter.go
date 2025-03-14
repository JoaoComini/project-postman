package memory

import (
	"github.com/OtavioPompolini/project-postman/internal/model"
)

type PersistanceAdapter interface {
	GetRequests() *map[int64]*model.Request
	CreateRequest(name string) *model.Request
	UpdateRequest(r *model.Request)
}
