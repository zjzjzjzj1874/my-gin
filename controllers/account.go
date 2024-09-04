package controllers

import (
	"net/http"

	"github/zjzjzjzj1874/my-gin/helper"
	"github/zjzjzjzj1874/my-gin/model"

	"github.com/gin-gonic/gin"
)

// ListAccounts godoc
// @Summary      List accounts
// @Description  get accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200  {array}   model.Account
// @Failure      400  {object}  helper.HTTPError
// @Failure      404  {object}  helper.HTTPError
// @Failure      500  {object}  helper.HTTPError
// @Router       /account [get]
func (c *Controller) ListAccounts(ctx *gin.Context) {
	//q := ctx.Request.URL.Query().Get("q")
	accounts, err := model.List()
	if err != nil {
		helper.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}
