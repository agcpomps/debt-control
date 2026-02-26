CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE clients (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    phone TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE loans (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    client_id UUID NOT NULL REFERENCES clients(id) ON DELETE CASCADE,

    principal NUMERIC(15,2) NOT NULL CHECK (principal > 0),

    -- 0.00 até 0.30 (0% até 30%)
    monthly_interest_rate NUMERIC(4,3) NOT NULL 
        CHECK (monthly_interest_rate >= 0 AND monthly_interest_rate <= 0.30),

    months INTEGER NOT NULL CHECK (months > 0),

    total_amount NUMERIC(15,2) NOT NULL,

    disbursed_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    loan_id UUID NOT NULL REFERENCES loans(id) ON DELETE CASCADE,
    amount NUMERIC(15,2) NOT NULL CHECK (amount > 0),
    paid_at TIMESTAMP NOT NULL DEFAULT now()
);