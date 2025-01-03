package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/my-go-app/pkg/model"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) CreateBrand(c *gin.Context) {
	var brand model.Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed"})
		return
	}

	if err := h.DB.Create(&brand).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, brand)
}

func (h *Handler) CreateVoucher(c *gin.Context) {
	var voucher model.Voucher
	if err := c.ShouldBindJSON(&voucher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(&voucher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed"})
		return
	}

	if err := h.DB.Create(&voucher).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, voucher)
}

func (h *Handler) GetVoucher(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid voucher id"})
		return
	}

	var voucher model.Voucher
	if err := h.DB.Preload("Brand").First(&voucher, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "voucher not found"})
		return
	}
	c.JSON(http.StatusOK, voucher)
}

func (h *Handler) GetVouchersByBrand(c *gin.Context) {
	brandID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid brand id"})
		return
	}

	var vouchers []model.Voucher
	if err := h.DB.Where("brand_id = ?", brandID).Preload("Brand").Find(&vouchers).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no vouchers found for this brand"})
		return
	}
	c.JSON(http.StatusOK, vouchers)
}

func (h *Handler) MakeRedemption(c *gin.Context) {
	var redemption model.Redemption
	if err := c.ShouldBindJSON(&redemption); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var totalPoints int
	for _, voucher := range redemption.Vouchers {
		totalPoints += int(voucher.CostInPoint)
	}

	// Ensure the total points is assigned as an int (not uint)
	redemption.TotalPoints = totalPoints
	if err := h.DB.Create(&redemption).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, redemption)
}

func (h *Handler) GetTransactionDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("transactionId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transaction id"})
		return
	}

	var redemption model.Redemption
	if err := h.DB.Preload("Vouchers").First(&redemption, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "transaction not found"})
		return
	}
	c.JSON(http.StatusOK, redemption)
}
