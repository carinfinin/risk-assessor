package model

import (
	"time"
)

type User struct {
	// Основные идентификаторы
	ID            string `json:"id" db:"id"`
	ApplicationID string `json:"application_id" db:"application_id"`
	ClientID      string `json:"client_id" db:"client_id"`

	// Шифрованные данные
	EncryptedClientData   []byte `json:"encrypted_client_data" db:"encrypted_client_data"`
	EncryptionKeyID       string `json:"encryption_key_id" db:"encryption_key_id"`
	EncryptedExternalData []byte `json:"encrypted_external_data" db:"encrypted_external_data"`

	// Хэши для поиска/верификации
	ClientIDHash string `json:"client_id_hash" db:"client_id_hash"`
	PassportHash string `json:"passport_hash" db:"passport_hash"`

	// Данные заявки
	LoanAmount float64 `json:"loan_amount" db:"loan_amount"`
	LoanTerm   int     `json:"loan_term" db:"loan_term"`

	// Статус и скоринг
	Status    string `json:"status" db:"status"`
	RiskScore int    `json:"risk_score" db:"risk_score"`

	// Решение
	DecisionReason string `json:"decision_reason" db:"decision_reason"`

	// Мета-данные
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ClientData struct {
	FullName       string `json:"full_name"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Income         int    `json:"income"`
	NumberPassport int    `json:"number_passport"`
	LoanAmount     int    `json:"loan_amount"`
	LoanTerm       int    `json:"loan_term"`
}
