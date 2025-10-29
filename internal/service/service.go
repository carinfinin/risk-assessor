package service

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/carinfinin/risk-assessor/internal/encryption"
	"github.com/carinfinin/risk-assessor/internal/model"
)

type Store interface {
	CreateUser(user *model.User) error
}
type Service struct {
	store     Store
	encryptor *encryption.Encryptor
}

func New( /*store Store, */ encryptor *encryption.Encryptor) *Service {
	return &Service{
		//store:     store,
		encryptor: encryptor,
	}
}

func (s *Service) CreateUser(clientData *model.ClientData) (*model.User, error) {

	passportHash := sha256.Sum256([]byte(clientData.NumberPassport))
	clientHash := sha256.Sum256([]byte(clientData.FullName + clientData.Phone))

	// проверка на дубль to do

	var user = model.User{
		ApplicationID: "",
		ClientID:      "",
		// Шифрованные данные
		EncryptedClientData:   nil,
		EncryptionKeyID:       "",
		EncryptedExternalData: nil,
		// Хэши для поиска/верификации
		ClientIDHash: hex.EncodeToString(clientHash[:]),
		PassportHash: hex.EncodeToString(passportHash[:]),
		// Данные заявки
		LoanAmount: clientData.LoanAmount,
		LoanTerm:   clientData.LoanTerm,
		// Статус и скоринг
		Status:    "new",
		RiskScore: 0,
		// Решение
		DecisionReason: "",
		//FullName
		//Phone
		//Email
		//Income
		//NumberPassport

	}
	err := s.store.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	return &model.User{}, nil
}
