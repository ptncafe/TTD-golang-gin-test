package update_shop_info

import (
	"TTD-golang-gin-test/common/constant"
	"TTD-golang-gin-test/dto"
	"TTD-golang-gin-test/entity"
	"TTD-golang-gin-test/interface/iproxy"
	"TTD-golang-gin-test/interface/irepository"
	mocksProxy "TTD-golang-gin-test/mocks/interface/iproxy"
	mocksRepo "TTD-golang-gin-test/mocks/interface/irepository"
	"time"

	"context"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"testing"
)

var phoneDefault = "0900000001"
var createdDateDefault = time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC)
var updatedDateDefault = time.Date(2021, time.Month(5), 21, 1, 10, 30, 0, time.UTC)

//https://github.com/stretchr/testify
//https://medium.com/@thegalang/testing-in-go-mocking-mvc-using-testify-and-mockery-c25344a88691
type UpdateNameTestSuite struct {
	suite.Suite
	ctx               context.Context
	log               *logrus.Logger
	updateNameRepo    irepository.IUpdateNameRepository
	getShopRepository irepository.IGetShopRepository
	pubSubProxy       proxy.IPubSubProxy
	shopEntityDefault entity.Shop
}

func TestUpdateName(t *testing.T) {
	suite.Run(t, new(UpdateNameTestSuite))
}
func (suite *UpdateNameTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.log = logrus.New()
	suite.updateNameRepo = new(mocksRepo.IUpdateNameRepository)
	suite.getShopRepository = new(mocksRepo.IGetShopRepository)
	suite.pubSubProxy = new(mocksProxy.IPubSubProxy)
	suite.shopEntityDefault = entity.Shop{Id: 1,
		Name:        "shop test name",
		Code:        "shop_test_name",
		Mobile:      &phoneDefault,
		Phone:       &phoneDefault,
		CreatedDate: &createdDateDefault,
		CreatedUser: phoneDefault,
		UpdatedDate: &updatedDateDefault,
		UpdatedUser: phoneDefault,
		SystemUrl:   "shop_test_name",
	}
}

func (suite *UpdateNameTestSuite) Test_When_Id_Is_Zero_Expect_BadRequest() {
	request := dto.UpdateNameRequest{Id: 0,
		Name:        "Test_When_Id_Is_Zero_Expect_BadRequest",
		UpdatedUser: "unit_test_user",
	}

	dom := NewDomain(suite.log,
		suite.updateNameRepo,
		suite.getShopRepository,
		suite.pubSubProxy,
	)

	err := dom.UpdateName(suite.ctx, request)
	suite.NotNil(err)
	suite.Error(err)
	if errors.IsBadRequest(err) == false {
		suite.Fail("Test_When_Id_Is_Zero_Expect_BadRequest IsBadRequest(err) == false")
	}
	suite.Errorf(err, constant.UpdateName_Error_Message_Id)

}

func (suite *UpdateNameTestSuite) Test_When_Name_Is_Too_Short_Expect_BadRequest() {
	request := dto.UpdateNameRequest{Id: 1, Name: "A"}

	dom := NewDomain(suite.log,
		suite.updateNameRepo,
		suite.getShopRepository,
		suite.pubSubProxy,
	)

	err := dom.UpdateName(suite.ctx, request)
	suite.NotNil(err)
	suite.Error(err)
	if errors.IsBadRequest(err) == false {
		suite.Fail("Test_When_Name_Is_Too_Short_Expect_BadRequest IsBadRequest(err) == false")
	}
	suite.Errorf(err, constant.UpdateName_Error_Message_Name_Short)

}

func (suite *UpdateNameTestSuite) Test_When_Name_Is_Too_Long_Expect_BadRequest() {
	request := dto.UpdateNameRequest{Id: 1,
		Name:        "Test_When_Name_Is_Too_Long_Expect_BadRequest_Test_When_Name_Is_Too_Long_Expect_BadRequest",
		UpdatedUser: "unit_test_user",
	}

	dom := NewDomain(suite.log,
		suite.updateNameRepo,
		suite.getShopRepository,
		suite.pubSubProxy,
	)

	err := dom.UpdateName(suite.ctx, request)
	suite.NotNil(err)
	suite.Error(err)
	if errors.IsBadRequest(err) == false {
		suite.Fail("Test_When_Name_Is_Too_Long_Expect_BadRequest IsBadRequest(err) == false")
	}
	suite.Errorf(err, constant.UpdateName_Error_Message_Name_Long)
}

func (suite *UpdateNameTestSuite) Test_When_UpdatedUser_Is_Nil_Expect_BadRequest() {
	request := dto.UpdateNameRequest{
		Id:          1,
		Name:        "Test_When_Name_Is_Too_Long_Expect_BadRequest",
		UpdatedUser: "",
	}
	mockGetShopRepo := new(mocksRepo.IGetShopRepository)
	mockGetShopRepo.On("GetById", suite.ctx, request.Id).Return(nil, nil)
	dom := NewDomain(suite.log,
		suite.updateNameRepo,
		mockGetShopRepo,
		suite.pubSubProxy,
	)
	err := dom.UpdateName(suite.ctx, request)
	suite.NotNil(err)
	suite.Error(err)
	if errors.IsBadRequest(err) == false {
		suite.Fail("Test_When_UpdatedUser_Is_Nil_Expect_BadRequest IsBadRequest(err) == false")
	}
	suite.Errorf(err, constant.UpdateName_Error_Message_UpdatedUser)
}

func (suite *UpdateNameTestSuite) Test_When_Shop_Is_Not_Exist_Expect_NotFound() {
	request := dto.UpdateNameRequest{Id: 1,
		Name:        "When_Shop_Not_Exist_Expect_NotFound",
		UpdatedUser: "unit_test_user",
	}

	mockGetShopRepo := new(mocksRepo.IGetShopRepository)
	mockGetShopRepo.On("GetById", suite.ctx, request.Id).Return(nil, nil)
	dom := NewDomain(suite.log,
		suite.updateNameRepo,
		mockGetShopRepo,
		suite.pubSubProxy,
	)
	err := dom.UpdateName(suite.ctx, request)
	suite.NotNil(err)
	suite.Error(err)
	if errors.IsNotFound(err) == false {
		suite.Fail("Test_When_Shop_Is_Not_Exist_Expect_NotFound IsNotFound(err) == false")
	}
	suite.Errorf(err, constant.UpdateName_Error_Message_Name_Long)
}

func (suite *UpdateNameTestSuite) Test_When_Name_Is_Duplicate_Expect_BadRequest() {
	request := dto.UpdateNameRequest{Id: 1,
		Name:        "Test_When_Name_Is_Duplicate_Expect_BadRequest",
		UpdatedUser: "unit_test_user",
	}
	mockGetShopRepo := new(mocksRepo.IGetShopRepository)
	mockGetShopRepo.On("GetById", suite.ctx, request.Id).Return(&suite.shopEntityDefault, nil)
	mockGetShopRepo.On("GetByName", suite.ctx, request.Name).Return(&suite.shopEntityDefault, nil)

	dom := NewDomain(suite.log,
		suite.updateNameRepo,
		mockGetShopRepo,
		suite.pubSubProxy,
	)
	err := dom.UpdateName(suite.ctx, request)
	suite.NotNil(err)
	suite.Error(err)
	if errors.IsBadRequest(err) == false {
		suite.Fail("Test_When_Name_Is_Duplicate_Expect_BadRequest IsBadRequest(err) == false")
	}
	suite.Errorf(err, constant.UpdateName_Error_Message_Name_Long)
}

func (suite *UpdateNameTestSuite) Test_Get_Id_Fail_Expect_Error() {
	request := dto.UpdateNameRequest{Id: 1,
		Name:        "Test_Get_Id_Fail_Expect_Error",
		UpdatedUser: "unit_test_user",
	}

	mockGetShopRepo := new(mocksRepo.IGetShopRepository)
	mockGetShopRepo.On("GetById", suite.ctx, request.Id).Return(nil, errors.New("Error"))


	dom := NewDomain(suite.log,
		suite.updateNameRepo,
		mockGetShopRepo,
		suite.pubSubProxy,
	)
	err := dom.UpdateName(suite.ctx, request)
	suite.NotNil(err)
	suite.Error(err)
}
func (suite *UpdateNameTestSuite) Test_Get_Name_Fail_Expect_Error() {
	request := dto.UpdateNameRequest{Id: 1,
		Name:        "Test_Get_Name_Fail_Expect_Error",
		UpdatedUser: "unit_test_user",
	}
	shopEntity := suite.shopEntityDefault

	mockGetShopRepo := new(mocksRepo.IGetShopRepository)
	mockGetShopRepo.On("GetById", suite.ctx, request.Id).Return(&shopEntity, nil)
	mockGetShopRepo.On("GetByName", suite.ctx, request.Name).Return(nil, errors.New("Error"))


	dom := NewDomain(suite.log,
		suite.updateNameRepo,
		mockGetShopRepo,
		suite.pubSubProxy,
	)
	err := dom.UpdateName(suite.ctx, request)
	suite.NotNil(err)
	suite.Error(err)
}



func (suite *UpdateNameTestSuite) Test_Update_DB_Fail_Expect_Error() {

	request := dto.UpdateNameRequest{Id: 1,
		Name:        "Test_Update_Success_Expect_Nil_Error",
		UpdatedUser: "unit_test_user",
	}
	shopEntity := suite.shopEntityDefault

	mockGetShopRepo := new(mocksRepo.IGetShopRepository)
	mockGetShopRepo.On("GetById", suite.ctx, request.Id).Return(&shopEntity, nil)
	mockGetShopRepo.On("GetByName", suite.ctx, request.Name).Return(nil, nil)

	shopEntity.Code = request.Code
	shopEntity.Name = request.Name
	shopEntity.UpdatedUser = request.UpdatedUser
	shopEntity.UpdatedDate = request.UpdatedDate

	mockUpdateNameRepo := new(mocksRepo.IUpdateNameRepository)
	mockUpdateNameRepo.On("UpdateName", suite.ctx, shopEntity).Return(errors.New("Error"))

	dom := NewDomain(suite.log,
		mockUpdateNameRepo,
		mockGetShopRepo,
		suite.pubSubProxy,
	)
	err := dom.UpdateName(suite.ctx, request)
	suite.NotNil(err)
	suite.Error(err)
}

func (suite *UpdateNameTestSuite) Test_Push_PubSub_Fail_Expect_Nil_Error() {
	request := dto.UpdateNameRequest{Id: 1,
		Name:        "Test_Update_Success_Expect_Nil_Error",
		UpdatedUser: "unit_test_user",
	}
	shopEntity := suite.shopEntityDefault
	mockGetShopRepo := new(mocksRepo.IGetShopRepository)
	mockGetShopRepo.On("GetById", suite.ctx, request.Id).Return(&shopEntity, nil)
	mockGetShopRepo.On("GetByName", suite.ctx, request.Name).Return(nil, nil)

	shopEntity.Code = request.Code
	shopEntity.Name = request.Name
	shopEntity.UpdatedUser = request.UpdatedUser
	shopEntity.UpdatedDate = request.UpdatedDate

	mockUpdateNameRepo := new(mocksRepo.IUpdateNameRepository)
	mockUpdateNameRepo.On("UpdateName", suite.ctx, shopEntity).Return(nil)

	mockPubSubProxy := new(mocksProxy.IPubSubProxy)
	mockPubSubProxy.On("PubShop", shopEntity).Return(errors.New("Error"))

	dom := NewDomain(suite.log,
		mockUpdateNameRepo,
		mockGetShopRepo,
		mockPubSubProxy,
	)
	err := dom.UpdateName(suite.ctx, request)
	suite.NotNil(err)
	suite.Error(err)
}

func (suite *UpdateNameTestSuite) Test_Update_DB_Success_Expect_Nil_Error() {
	request := dto.UpdateNameRequest{Id: 1,
		Name:        "Test_Update_Success_Expect_Nil_Error",
		UpdatedUser: "unit_test_user",
	}
	shopEntity := suite.shopEntityDefault

	mockGetShopRepo := new(mocksRepo.IGetShopRepository)
	mockGetShopRepo.On("GetById", suite.ctx, request.Id).Return(&shopEntity, nil)
	mockGetShopRepo.On("GetByName", suite.ctx, request.Name).Return(nil, nil)

	shopEntity.Code = request.Code
	shopEntity.Name = request.Name
	shopEntity.UpdatedUser = request.UpdatedUser
	shopEntity.UpdatedDate = request.UpdatedDate

	mockUpdateNameRepo := new(mocksRepo.IUpdateNameRepository)
	mockUpdateNameRepo.On("UpdateName", suite.ctx, shopEntity).Return(nil)

	mockPubSubProxy := new(mocksProxy.IPubSubProxy)
	mockPubSubProxy.On("PubShop", shopEntity).Return(nil)

	dom := NewDomain(suite.log,
		mockUpdateNameRepo,
		mockGetShopRepo,
		mockPubSubProxy,
	)
	err := dom.UpdateName(suite.ctx, request)
	suite.Nil(err)
	suite.NoError(err)
}



