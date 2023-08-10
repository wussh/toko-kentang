package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/wussh/tokokentang/features/transaction"

	"github.com/labstack/echo/v4"
)

type trxHandle struct {
	srv transaction.TrxService
}

func New(ts transaction.TrxService) transaction.TrxHandler {
	return &trxHandle{
		srv: ts,
	}
}

func (th *trxHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {

		input := AddTrxRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "invalid input")
		}

		cartID, err := strconv.Atoi(c.Param("id"))

		cnv := ToCore(input)

		err = th.srv.Add(uint(cartID), c.Get("user"), *cnv)
		if err != nil {
			log.Println("error post content : ", err.Error())
			return c.JSON(http.StatusInternalServerError, "unable to process the data")
		}
		return c.JSON(http.StatusCreated, "Created a new transactions succesfully")
	}
}
