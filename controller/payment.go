package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/zakariawahyu/go-gin-gorm-mvc/entity"
	"github.com/zakariawahyu/go-gin-gorm-mvc/models"
	"net/http"
	"strconv"
)

func CreatePayment(c *gin.Context) {
	var payment entity.Payment
	var order entity.OrderResponse
	if err := c.ShouldBind(&payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := models.GetOrderById(&order, payment.OrderID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var grossAmount int
	itemDetail := make([]midtrans.ItemDetails, 0)

	for _, detailOrder := range order.OrderDetail {
		grossAmount += detailOrder.Qty * detailOrder.Product.Price
		itemDetail = append(itemDetail, midtrans.ItemDetails{
			ID:    strconv.Itoa(detailOrder.ProductID),
			Name:  detailOrder.Product.NameProduct,
			Price: int64(detailOrder.Product.Price),
			Qty:   int32(detailOrder.Qty),
		})
	}

	req := &coreapi.ChargeReq{
		PaymentType: "bank_transfer",
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(order.ID),
			GrossAmt: int64(grossAmount),
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: order.Customer.FullName,
			Email: order.Customer.Email,
			Phone: order.Customer.PhoneNumber,
		},
		Items: &itemDetail,
	}

	midtrans.ServerKey = "SB-Mid-server-ONEEAiYCW7UYfw6nstBB2iRK"
	midtrans.Environment = midtrans.Sandbox
	result, _ := coreapi.ChargeTransaction(req)
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
