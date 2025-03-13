package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"test-echo/internal/dto"
	"test-echo/internal/service"
)

type NasabahHandler struct {
	service service.NasabahService
}

func NewNasabahHandler(service service.NasabahService) *NasabahHandler {
	return &NasabahHandler{service: service}
}

func (h *NasabahHandler) DaftarNasabah(c echo.Context) error {
	var req dto.NasabahRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Remark: "Invalid request payload"})
	}

	noRekening, err := h.service.DaftarNasabah(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Remark: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.NasabahResponse{NoRekening: noRekening})
}

func (h *NasabahHandler) Tabung(c echo.Context) error {
	var req dto.TabungRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Remark: "Invalid request payload"})
	}

	saldo, err := h.service.Tabung(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Remark: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.TabungResponse{Saldo: saldo})
}

func (h *NasabahHandler) Tarik(c echo.Context) error {
	var req dto.TarikRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Remark: "Invalid request payload"})
	}

	saldo, err := h.service.Tarik(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Remark: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.TarikResponse{Saldo: saldo})
}

func (h *NasabahHandler) GetSaldo(c echo.Context) error {
	noRekening := c.Param("no_rekening")

	saldo, err := h.service.GetSaldo(noRekening)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Remark: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SaldoResponse{Saldo: saldo})
}