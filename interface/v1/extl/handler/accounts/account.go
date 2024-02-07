package accounts

import (
	"dealls/core/v1/port/account"
	"github.com/labstack/echo/v4"
	"net/http"

	"dealls/core"
	"dealls/interface/v1/general/common"
	"dealls/pkg/util"
)

var log = util.NewLogger()

type Handler struct {
	accountUsecase account.AccountUsecase
}

func New(service account.AccountUsecase) *Handler {
	return &Handler{
		accountUsecase: service,
	}
}

// List 	godoc
// @Summary		List profiles.
// @Description	List profiles.
// @Tags		Accounts
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string		true	"Bearer token"
// @Param		page			query		string		false 	"1"
// @Param		limit			query		string		false 	"10"
// @Failure		400
// @Failure		401
// @Failure		422
// @Failure		500
// @Router	/api/v1/accounts [get]
func (h *Handler) List(c echo.Context) error {
	ic := c.Get("ic").(*core.InternalContext)
	claims := c.Get("claims").(common.JwtClaims)

	request := new(List)
	if err := c.Bind(request); err != nil {
		res := common.HandleBind(err)
		return c.JSON(res.StatusCode, common.NewResponse(res, nil))
	}

	if err := c.Validate(request); err != nil {
		res := common.ResultMap[core.BAD_REQUEST]
		res.Errors = util.GetValidatorMessage(err)
		return c.JSON(http.StatusBadRequest, common.NewResponse(res, nil))
	}

	if request.Page < 1 || request.Limit < 1 {
		request.Page = 1
		request.Limit = 10
	}

	result, total, cerr := h.accountUsecase.AccountList(ic, claims.Id, map[string]any{
		"page":  request.Page,
		"limit": request.Limit,
	})
	if cerr != nil {
		res := common.HandleError(cerr)
		return c.JSON(res.StatusCode, common.NewResponse(res, nil))
	}

	return c.JSON(
		http.StatusOK,
		common.NewListResponseWithMeta(common.ResultMap[core.OK], result, common.GetMeta(request.Page, request.Limit, total)),
	)
}

// Get 	godoc
// @Summary		Get profile.
// @Description	Get profile.
// @Tags		Accounts
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string		true	"Bearer token"
// @Param		id				path		string		false 	"65c1de91056ae9755c64ffba"
// @Failure		400
// @Failure		401
// @Failure		422
// @Failure		500
// @Router	/api/v1/accounts/{id} [get]
func (h *Handler) Get(c echo.Context) error {
	ic := c.Get("ic").(*core.InternalContext)

	request := new(Get)
	if err := c.Bind(request); err != nil {
		res := common.HandleBind(err)
		return c.JSON(res.StatusCode, common.NewResponse(res, nil))
	}

	if err := c.Validate(request); err != nil {
		res := common.ResultMap[core.BAD_REQUEST]
		res.Errors = util.GetValidatorMessage(err)
		return c.JSON(http.StatusBadRequest, common.NewResponse(res, nil))
	}

	result, cerr := h.accountUsecase.AccountGet(ic, request.Id)
	if cerr != nil {
		res := common.HandleError(cerr)
		return c.JSON(res.StatusCode, common.NewResponse(res, nil))
	}

	return c.JSON(
		http.StatusOK,
		common.NewResponse(common.ResultMap[core.OK], result),
	)
}

// Action 	godoc
// @Summary		Action profile.
// @Description	Action .
// @Tags		Accounts
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string	true	"Bearer token"
// @Param  		Body 			body		Action 	true	"body"
// @Failure		400
// @Failure		401
// @Failure		422
// @Failure		500
// @Router	/api/v1/accounts/{id}/action [post]
func (h *Handler) Action(c echo.Context) error {
	ic := c.Get("ic").(*core.InternalContext)
	claims := c.Get("claims").(common.JwtClaims)

	request := new(Action)
	if err := c.Bind(request); err != nil {
		res := common.HandleBind(err)
		return c.JSON(res.StatusCode, common.NewResponse(res, nil))
	}

	if err := c.Validate(request); err != nil {
		res := common.ResultMap[core.BAD_REQUEST]
		res.Errors = util.GetValidatorMessage(err)
		return c.JSON(http.StatusBadRequest, common.NewResponse(res, nil))
	}

	cerr := h.accountUsecase.AccountAction(ic, claims.Id, request.Id, request.Action)
	if cerr != nil {
		res := common.HandleError(cerr)
		return c.JSON(res.StatusCode, common.NewResponse(res, nil))
	}

	return c.JSON(
		http.StatusOK,
		common.NewResponse(common.ResultMap[core.OK], nil),
	)
}

// Activate 	godoc
// @Summary		Activate premium package.
// @Description	Activate premium package.
// @Tags		Accounts
// @Accept		json
// @Produce		json
// @Param		Authorization	header		string	true	"Bearer token"
// @Param		id				path		string		false 	"65c1de91056ae9755c64ffba"
// @Failure		400
// @Failure		401
// @Failure		422
// @Failure		500
// @Router	/api/v1/accounts/{id}/activate [post]
func (h *Handler) Activate(c echo.Context) error {
	ic := c.Get("ic").(*core.InternalContext)

	request := new(Premium)
	if err := c.Bind(request); err != nil {
		res := common.HandleBind(err)
		return c.JSON(res.StatusCode, common.NewResponse(res, nil))
	}

	if err := c.Validate(request); err != nil {
		res := common.ResultMap[core.BAD_REQUEST]
		res.Errors = util.GetValidatorMessage(err)
		return c.JSON(http.StatusBadRequest, common.NewResponse(res, nil))
	}

	result, cerr := h.accountUsecase.AccountActivate(ic, request.Id)
	if cerr != nil {
		res := common.HandleError(cerr)
		return c.JSON(res.StatusCode, common.NewResponse(res, nil))
	}

	return c.JSON(
		http.StatusCreated,
		common.NewResponse(common.ResultMap[core.CREATED], result),
	)
}
