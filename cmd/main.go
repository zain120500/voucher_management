package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/my-go-app/config"
	"github.com/my-go-app/pkg/handler"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	r := gin.Default()

	h := handler.NewHandler(db)

	r.POST("/brand", h.CreateBrand)
	r.POST("/voucher", h.CreateVoucher)
	r.GET("/voucher", h.GetVoucher)
	r.GET("/voucher/brand", h.GetVouchersByBrand)
	r.POST("/transaction/redemption", h.MakeRedemption)
	r.GET("/transaction/redemption", h.GetTransactionDetail)

	r.Run(":8080")
}
