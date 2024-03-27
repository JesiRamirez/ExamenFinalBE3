package handler

import (
	"strconv"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/dentist"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
	"github.com/gin-gonic/gin"
)

type dentistHandler struct {
	s dentist.Service
}

// NewDentistHandler create a new dentist controller
func NewDentistHandler(s dentist.Service) *dentistHandler {
	return &dentistHandler{
		s: s,
	}
}

// POST Create a new dentist
func (h *dentistHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dentist domain.Dentist
		err := ctx.ShouldBindJSON(&dentist)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid dentist"})
			return
		}

		p, err := h.s.Create(dentist)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, p)
	}
}

// GET all dentist
func (h *dentistHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		dentists, _ := h.s.GetAll()
		c.JSON(200, dentists)
	}
}

// GET dentist by ID
func (h *dentistHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		dentist, err := h.s.GetByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "dentist not found"})
			return
		}
		c.JSON(200, dentist)
	}
}

// UPDATE dentist
func (h *dentistHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idString := ctx.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}

		var dentist domain.Dentist
		err = ctx.ShouldBindJSON(&dentist)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid dentist"})
			return
		}

		p, err := h.s.Update(id, dentist)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(201, p)
	}

}

// Patch dentist
func (h *dentistHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name     string `json:"name,omitempty"`
		Lastname string `json:"lastname,omitempty"`
		License  string `json:"license,omitempty"`
	}

	return func(ctx *gin.Context) {
		var r Request
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(400, gin.H{"error": "invalid json"})
			return
		}
		update := domain.Dentist{
			Name:     r.Name,
			Lastname: r.Lastname,
			License:  r.License,
		}
		p, err := h.s.Patch(id, update)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, p)

	}
}

// Delete elimina un producto
func (h *dentistHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(204, gin.H{"msg": "product deleted"})
	}
}
