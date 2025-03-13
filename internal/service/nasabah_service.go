package service

import (
	"errors"

	"test-echo/internal/dto"
	"test-echo/internal/repository"
)

type NasabahService interface {
	DaftarNasabah(req dto.NasabahRequest) (string, error)
	Tabung(req dto.TabungRequest) (float64, error)
	Tarik(req dto.TarikRequest) (float64, error)
	GetSaldo(noRekening string) (float64, error)
}

type nasabahService struct {
	repo repository.NasabahRepository
}

func NewNasabahService(repo repository.NasabahRepository) NasabahService {
	return &nasabahService{repo: repo}
}

func (s *nasabahService) DaftarNasabah(req dto.NasabahRequest) (string, error) {
	if s.repo.IsNIKExists(req.NIK) {
		return "", errors.New("NIK already exists")
	}
	if s.repo.IsNoHPExists(req.NoHP) {
		return "", errors.New("NoHP already exists")
	}

	noRekening, err := s.repo.CreateNasabah(req)
	if err != nil {
		return "", err
	}

	return noRekening, nil
}

func (s *nasabahService) Tabung(req dto.TabungRequest) (float64, error) {
	// Cek apakah no_rekening valid
	saldo, err := s.repo.GetSaldo(req.NoRekening)
	if err != nil {
		return 0, errors.New("no_rekening tidak dikenali")
	}

	// Lakukan transaksi tabung
	if err := s.repo.Tabung(req.NoRekening, req.Nominal); err != nil {
		return 0, err
	}

	// Ambil saldo terbaru
	saldo, err = s.repo.GetSaldo(req.NoRekening)
	if err != nil {
		return 0, err
	}

	return saldo, nil
}

func (s *nasabahService) Tarik(req dto.TarikRequest) (float64, error) {
	// Cek apakah no_rekening valid
	saldo, err := s.repo.GetSaldo(req.NoRekening)
	if err != nil {
		return 0, errors.New("no_rekening tidak dikenali")
	}

	// Lakukan transaksi tarik
	if err := s.repo.Tarik(req.NoRekening, req.Nominal); err != nil {
		return 0, err
	}

	// Ambil saldo terbaru
	saldo, err = s.repo.GetSaldo(req.NoRekening)
	if err != nil {
		return 0, err
	}

	return saldo, nil
}

func (s *nasabahService) GetSaldo(noRekening string) (float64, error) {
	saldo, err := s.repo.GetSaldo(noRekening)
	if err != nil {
		return 0, errors.New("no_rekening tidak dikenali")
	}
	return saldo, nil
}