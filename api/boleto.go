package api

import (
	"net/http"

	"time"

	"bitbucket.org/mundipagg/boletoapi/bank"
	"bitbucket.org/mundipagg/boletoapi/boleto"
	"bitbucket.org/mundipagg/boletoapi/models"
	gin "gopkg.in/gin-gonic/gin.v1"
)

//Regista um boleto em um determinado banco
func registerBoleto(c *gin.Context) {
	boleto := models.BoletoRequest{}
	errBind := c.BindJSON(&boleto)
	//TODO melhorar isso
	if errBind != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Code: "000", Message: errBind.Error()})
		return
	}

	d, errFmt := time.Parse("2006-01-02", boleto.Title.ExpireDate)
	boleto.Title.ExpireDateTime = d
	if errFmt != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Code: "000", Message: errFmt.Error()})
		return
	}
	bank, err := bank.Get(boleto.BankNumber)

	lg := bank.Log()
	lg.Operation = "RegisterBoleto"
	lg.NossoNumero = boleto.Title.OurNumber

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Code: "001", Message: err.Error()})
		return
	}

	lg.Recipient = bank.GetBankNumber().BankName()
	lg.Request(boleto, c.Request.URL.RequestURI(), c.Request.Header)

	resp, errR := bank.RegisterBoleto(boleto)
	if errR != nil {
		errResp := models.BoletoResponse{
			Errors: models.NewSingleErrorCollection("MP400", err.Error()),
		}
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	lg.Response(resp, c.Request.URL.RequestURI())
	st := http.StatusOK
	if len(resp.Errors) > 0 {
		st = http.StatusBadRequest
	}
	c.JSON(st, resp)
}

func getBoleto(c *gin.Context) {
	c.Status(200)
	c.Header("Content-Type", "text/html")
	boleto.HTML(c.Writer, nil)
}
