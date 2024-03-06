package services

import (
	"framework-iris/pojo/request"
	"framework-iris/pojo/response"
	"github.com/qiafan666/gotato/commons"
	"sync"
)

// BaseService service layer interface
type BaseService interface {
	Test(info request.Test) (out response.Test, code commons.ResponseCode, err error)
}

var baseServiceIns *baseServiceImp
var baseServiceInitOnce sync.Once

func NewBaseServiceInstance() BaseService {

	baseServiceInitOnce.Do(func() {
		baseServiceIns = &baseServiceImp{
			//dao: dao.Instance(),
		}
	})

	return baseServiceIns
}

type baseServiceImp struct {
	//dao dao.Dao
}

func (p baseServiceImp) Test(info request.Test) (out response.Test, code commons.ResponseCode, err error) {
	return
}
