package dto

type NasabahRequest struct {
	Nama string `json:"nama"`
	NIK  string `json:"nik"`
	NoHP string `json:"no_hp"`
}

type NasabahResponse struct {
	NoRekening string `json:"no_rekening"`
}

type ErrorResponse struct {
	Remark string `json:"remark"`
}

type Nasabah struct {
	NoRekening string `gorm:"primaryKey"`
	Nama       string
	NIK        string `gorm:"unique"`
	NoHP       string `gorm:"unique"`
	Saldo      float64
}

type TabungRequest struct {
	NoRekening string  `json:"no_rekening"`
	Nominal    float64 `json:"nominal"`
}

type TabungResponse struct {
	Saldo float64 `json:"saldo"`
}

type TarikRequest struct {
	NoRekening string  `json:"no_rekening"`
	Nominal    float64 `json:"nominal"`
}

type TarikResponse struct {
	Saldo float64 `json:"saldo"`
}

type SaldoResponse struct {
	Saldo float64 `json:"saldo"`
}