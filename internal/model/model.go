package model

// 1. НБКИ (Национальное бюро кредитных историй)
type NBKIResponse struct {
	CreditScore    int  `json:"credit_score"`    // 300-850
	ActiveLoans    int  `json:"active_loans"`    // текущие кредиты
	OverdueLoans   int  `json:"overdue_loans"`   // просрочки
	DefaultHistory bool `json:"default_history"` // были ли дефолты
	InquiryCount   int  `json:"inquiry_count"`   // кол-во запросов
}

// 2. ФССП (Федеральная служба судебных приставов)
type FSSPResponse struct {
	HasEnforcementProceedings bool     `json:"has_enforcement_proceedings"`
	DebtAmount                *float64 `json:"debt_amount,omitempty"`
	CasesCount                int      `json:"cases_count"`
}

// 3. МВД (проверка паспорта)
type MVDResponse struct {
	IsPassportValid  bool   `json:"is_passport_valid"`
	VerificationDate string `json:"verification_date"`
}

// 4. Прочие источники
type OtherSources struct {
	PEPCheck   bool `json:"pep_check"`  // Politically Exposed Person
	Sanctions  bool `json:"sanctions"`  // в санкционных списках
	Bankruptcy bool `json:"bankruptcy"` // банкротство
}
