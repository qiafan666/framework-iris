package services

import (
	"framework-go/dao"
	"framework-go/pojo/request"
	"framework-go/pojo/response"
	"github.com/qiafan666/quickweb/commons"
	"sync"
)

// TestService service layer interface
type TestService interface {
	Test(info request.Test) (out response.Test, code commons.ResponseCode, err error)
}

var testServiceIns *testServiceImp
var testServiceInitOnce sync.Once

func NewTestServiceInstance() TestService {

	testServiceInitOnce.Do(func() {
		testServiceIns = &testServiceImp{
			dao: dao.Instance(),
		}
	})

	return testServiceIns
}

type testServiceImp struct {
	dao dao.Dao
}

func (p testServiceImp) Test(info request.Test) (out response.Test, code commons.ResponseCode, err error) {

	return
}
