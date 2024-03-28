package handler

import (
	"errors"
	"strconv"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/dentist"
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
	"github.com/bootcamp-go/ExamenFinalBE3.git/pkg/web"
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
// Post godoc
// @Summary      Create a new dentist
// @Description  Create a new dentist in repository
// @Tags         dentists
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Dentist true "Dentist"
// @Success      201 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /dentist [post]
func (h *dentistHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dentist domain.Dentist
		err := ctx.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(ctx, 400, errors.New("invalid dentist"))
			return
		}

		p, err := h.s.Create(dentist)
		if err != nil {
			web.Failure(ctx, 400, err)
			return
		}
		web.Success(ctx, 201, p)
	}
}

// GET all dentist
// GetAll godoc
// @Summary      Gets all the dentists
// @Description  Gets all the dentists from the repository
// @Tags         dentists
// @Produce      json
// @Success      200 {object}  web.response
// @Router       /dentists [get]
func (h *dentistHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		dentists, _ := h.s.GetAll()
		c.JSON(200, dentists)
	}
}

// GET dentist by ID
// GetByID godoc
// @Summary      Gets a dentist by id
// @Description  Gets a dentist by id from the repository
// @Tags         dentists
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /dentists/{id} [get]
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
// Put godoc
// @Summary      Updates a dentist
// @Description  Updates a dentist from the repository
// @Tags         dentists
// @Produce      json
// @Param        token header string true "token"
// @Param        id path string true "ID"
// @Param        body body domain.Dentist true "Dentist"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /dentists/{id} [put]
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

// PATCH dentist
// Patch godoc
// @Summary      Updates selected fields
// @Description  Updates selected fields from a dentist from the repository
// @Tags         dentists
// @Produce      json
// @Param        token header string true "token"
// @Param        id path string true "ID"
// @Param        body body domain.Dentist true "Dentist"
// @Success      200 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Router       /dentists/{id} [patch]
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

// DELETE elimina un dentista
// Delete godoc
// @Summary      Deletes a dentist
// @Description  Deletes a dentist from the repository
// @Tags         dentists
// @Produce      json
// @Param        token header string true "token"
// @Param        id path string true "ID"
// @Success      204 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /dentists/{id} [delete]
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
