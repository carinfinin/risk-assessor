CREATE TYPE users_status AS ENUM (
    'new', 'in_progress', 'approved', 'denied', 'manual_review'
);

CREATE TABLE users (
      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
      application_id VARCHAR(255) UNIQUE NOT NULL,
      client_id VARCHAR(255) NOT NULL,

    -- Зашифрованные данные
      encrypted_client_data BYTEA NOT NULL,
      encryption_key_id VARCHAR(100) NOT NULL,

    -- Мета-данные для индексации (опционально, хэшированные)
      client_id_hash VARCHAR(64), -- SHA256 от client_id для поиска
      passport_hash VARCHAR(64),  -- SHA256 от паспорта для дедубликации

      loan_amount DECIMAL(15,2) NOT NULL,
      loan_term INTEGER NOT NULL,
      status users_status NOT NULL DEFAULT 'new',
      risk_score INTEGER,

    -- Внешние данные тоже шифруем
      encrypted_external_data BYTEA, -- данные из внешних апи

      decision_reason TEXT, -- почему отказ или разреён
      created_at TIMESTAMPTZ DEFAULT NOW(),
      updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Индексы на хэши
CREATE INDEX idx_users_client_hash ON users(client_id_hash);
CREATE INDEX idx_users_passport_hash ON users(passport_hash);