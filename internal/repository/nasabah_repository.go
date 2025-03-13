package repository

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"

	"test-echo/internal/dto"
)

type NasabahRepository interface {
	IsNIKExists(nik string) bool
	IsNoHPExists(noHP string) bool
	CreateNasabah(req dto.NasabahRequest) (string, error)
	GetSaldo(noRekening string) (float64, error)
	Tabung(noRekening string, nominal float64) error
	Tarik(noRekening string, nominal float64) error
}

type nasabahRepository struct {
	db *gorm.DB
}

func NewNasabahRepository(db *gorm.DB) NasabahRepository {
	return &nasabahRepository{db: db}
}

func (r *nasabahRepository) IsNIKExists(nik string) bool {
	var count int64
	r.db.Model(&dto.Nasabah{}).Where("nik = ?", nik).Count(&count)
	return count > 0
}

func (r *nasabahRepository) IsNoHPExists(noHP string) bool {
	var count int64
	r.db.Model(&dto.Nasabah{}).Where("no_hp = ?", noHP).Count(&count)
	return count > 0
}

func (r *nasabahRepository) CreateNasabah(req dto.NasabahRequest) (string, error) {

	noRekening := generateNoRekening()

	nasabah := dto.Nasabah{
		Nama:  req.Nama,
		NIK:   req.NIK,
		NoHP:  req.NoHP,
		NoRekening: noRekening,
		Saldo: 0, // Default saldo awal
	}

	fmt.Printf("Nasabah Data: %+v\n", nasabah)

	if err := r.db.Create(&nasabah).Error; err != nil {
		return "", err
	}

	return nasabah.NoRekening, nil
}

func generateNoRekening() string {
	// Get current date in YYMMDD format
	today := time.Now().Format("060102") // YYMMDD format

	// Generate a random 6-digit number
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(900000) + 100000 // Ensures a 6-digit number (100000-999999)

	// Combine date and random number
	return fmt.Sprintf("%s%d", today, randomNumber)
}

func (r *nasabahRepository) GetSaldo(noRekening string) (float64, error) {
	var nasabah dto.Nasabah
	if err := r.db.Where("no_rekening = ?", noRekening).First(&nasabah).Error; err != nil {
		return 0, err
	}
	return nasabah.Saldo, nil
}

func (r *nasabahRepository) Tabung(noRekening string, nominal float64) error {
	return r.db.Model(&dto.Nasabah{}).Where("no_rekening = ?", noRekening).
		Update("saldo", gorm.Expr("saldo + ?", nominal)).Error
}

func (r *nasabahRepository) Tarik(noRekening string, nominal float64) error {
	// Cek saldo nasabah
	saldo, err := r.GetSaldo(noRekening)
	if err != nil {
		return errors.New("no_rekening tidak dikenali")
	}

	// Cek apakah saldo cukup
	if saldo < nominal {
		return errors.New("saldo tidak cukup")
	}

	// Lakukan transaksi tarik
	return r.db.Model(&dto.Nasabah{}).Where("no_rekening = ?", noRekening).
		Update("saldo", gorm.Expr("saldo - ?", nominal)).Error
}