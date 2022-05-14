package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mikailyusuf/go/test/internal"
)

type postProduct struct {
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Description string  `json:"description" binding:"omitempty,max=250"`
}

type Product struct {
	GUID        string  `json:"uuid"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CreatedAt   string  `json:createdAt`
}

func PostProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload postProduct
		var ctx = c.Request.Context()

		if e := c.ShouldBindJSON(&payload); e != nil {
			c.JSON(http.StatusBadRequest, internal.NewHTTPResponse(http.StatusBadRequest, e))
			return
		}

		var guuid = uuid.New().String()
		var createdAt = time.Now().Format(time.RFC3339)
		_, e := db.ExecContext(ctx,
			"INSERT INTO products (guid,name,price,description,createdAt) VALUES (?,?,?,?,?)", guuid, payload.Name, payload.Price, payload.Description, createdAt)
		if e != nil {
			var res = internal.NewHTTPResponse(http.StatusInternalServerError, e)
			c.JSON(http.StatusInternalServerError, res)
			return
		}
		var product Product
		var row = db.QueryRow("SELECT guid,name,price,description,createdAt FROM products WHERE guid=?", guuid)
		if e := row.Scan(&product.GUID, &product.Name, &product.Price, &product.Description, &product.CreatedAt); e != nil {
			var res = internal.NewHTTPResponse(http.StatusInternalServerError, e)
			c.JSON(http.StatusInternalServerError, res)
			return
		}

		var res = internal.NewHTTPResponse(http.StatusCreated, product)

		c.Writer.Header().Add("Location", fmt.Sprintf("/products/%s", guuid))
		c.JSON(http.StatusCreated, res)

		fmt.Println(payload)
	}
}
