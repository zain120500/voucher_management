package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/my-go-app/pkg/handler"
	"github.com/my-go-app/pkg/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&model.Brand{}, &model.Voucher{}, &model.Redemption{})
	return db
}

func setupRouter(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.POST("/brands", h.CreateBrand)
	r.POST("/vouchers", h.CreateVoucher)
	r.GET("/vouchers", h.GetVoucher)
	r.GET("/vouchers/by-brand", h.GetVouchersByBrand)
	r.POST("/redemptions", h.MakeRedemption)
	r.GET("/transactions", h.GetTransactionDetail)

	return r
}

func TestCreateBrand(t *testing.T) {
	db := setupTestDB()
	h := handler.NewHandler(db)
	r := setupRouter(h)

	brand := model.Brand{Name: "Test Brand"}
	jsonValue, _ := json.Marshal(brand)

	req, _ := http.NewRequest(http.MethodPost, "/brands", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdBrand model.Brand
	db.First(&createdBrand)
	assert.Equal(t, brand.Name, createdBrand.Name)
}

func TestCreateVoucher(t *testing.T) {
	db := setupTestDB()
	h := handler.NewHandler(db)
	r := setupRouter(h)

	brand := model.Brand{Name: "Test Brand"}
	db.Create(&brand)

	voucher := model.Voucher{Title: "Test Voucher", CostInPoint: 100, BrandID: brand.ID}
	jsonValue, _ := json.Marshal(voucher)

	req, _ := http.NewRequest(http.MethodPost, "/vouchers", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdVoucher model.Voucher
	db.First(&createdVoucher)
	assert.Equal(t, voucher.Title, createdVoucher.Title)
}

func TestGetVoucher(t *testing.T) {
	db := setupTestDB()
	h := handler.NewHandler(db)
	r := setupRouter(h)

	brand := model.Brand{Name: "Test Brand"}
	db.Create(&brand)

	voucher := model.Voucher{Title: "Test Voucher", CostInPoint: 100, BrandID: brand.ID}
	db.Create(&voucher)

	req, _ := http.NewRequest(http.MethodGet, "/vouchers?id="+strconv.Itoa(int(voucher.ID)), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseVoucher model.Voucher
	json.Unmarshal(w.Body.Bytes(), &responseVoucher)
	assert.Equal(t, voucher.Title, responseVoucher.Title)
}

func TestGetVouchersByBrand(t *testing.T) {
	db := setupTestDB()
	h := handler.NewHandler(db)
	r := setupRouter(h)

	brand := model.Brand{Name: "Test Brand"}
	db.Create(&brand)

	voucher := model.Voucher{Title: "Test Voucher", CostInPoint: 100, BrandID: brand.ID}
	db.Create(&voucher)

	req, _ := http.NewRequest(http.MethodGet, "/vouchers/by-brand?id="+strconv.Itoa(int(brand.ID)), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseVouchers []model.Voucher
	json.Unmarshal(w.Body.Bytes(), &responseVouchers)
	assert.Len(t, responseVouchers, 1)
	assert.Equal(t, voucher.Title, responseVouchers[0].Title)
}