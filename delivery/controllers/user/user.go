package user

import (
	"net/http"

	_common "sirclo/graphql/delivery/common"
	_userRepo "sirclo/graphql/repository/user"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repo _userRepo.User
}

func New(repository _userRepo.User) *UserController {
	return &UserController{repo: repository}
}

func (pc UserController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := pc.repo.Get()
		if err != nil {
			// return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			// 	"message": "Error on server",
			// 	"data":    nil,
			// })
			return _common.NewErrorResponse(c, http.StatusInternalServerError, "failed to get all users")
		}
		// return c.JSON(http.StatusOK, map[string]interface{}{
		// 	"message": "success get data",
		// 	"data":    res,
		// })
		var userResponseData []UserResponseFormat
		for _, v := range res {
			userResponseData = append(userResponseData, FromEntity(v))
		}
		return _common.NewSuccessResponse(c, "success get users", userResponseData)
	}
}

// func (pc UserController) Create() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		req := UserRequestFormat{}
// 		if err := c.Bind(&req); err != nil {
// 			// return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			// 	"message": "wrong input format",
// 			// 	"data":    req,
// 			// })
// 			return _common.NewErrorResponse(c, http.StatusBadRequest, "failed to bind data")
// 		}

// 		res, err := pc.repo.Create(*req.ToEntity())

// 		if err != nil {
// 			// return c.JSON(http.StatusBadGateway, map[string]interface{}{
// 			// 	"message": "theres is some problem",
// 			// 	"data":    req,
// 			// })
// 			return _common.NewErrorResponse(c, http.StatusInternalServerError, "failed to get all users")
// 		}

// 		// return c.JSON(http.StatusCreated, map[string]interface{}{
// 		// 	"message": "success register",
// 		// 	"data":    res,
// 		// })
// 		return _common.NewSuccessResponse(c, "success get data", res)
// 	}
// }
