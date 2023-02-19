package common

import (
	"fmt"
	"marketplace/pkg/utils"
	"time"
)

type BaseObject struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Hash      string    `json:"hash"`
}

func NewBaseObject() *BaseObject {
	base := &BaseObject{
		ID:        utils.GenerateUint64ID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Time{},
	}

	base.Hash = utils.GenerateSHA256Hash(fmt.Sprintf("%v", base))
	return base
}

func (base *BaseObject) BaseObjectUpdated() {
	base.UpdatedAt = time.Now()
	base.Hash = utils.GenerateSHA256Hash(fmt.Sprintf("%v", base))
	
}

func (base *BaseObject) BaseObjectDeleted() {
	base.DeletedAt = time.Now()
	base.Hash = utils.GenerateSHA256Hash(fmt.Sprintf("%v", base))
}
