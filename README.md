# Fintech Wallet Service

A backend wallet and payment system written in Go, designed using real-world fintech principles such as double-entry bookkeeping, immutable ledgers, and transaction-safe workflows.

This project is not a toy CRUD API. It models how money actually moves in production systems, focusing on correctness, auditability, and concurrency safety.

---

## Core Concepts

- **Immutable Ledger**  
  All monetary changes are recorded as append-only ledger entries. Balances are derived, never stored.

- **Double-Entry Accounting**  
  Every transaction produces one or more ledger entries that always balance.

- **Transactional Safety**  
  All money movements occur inside database transactions with row-level locking to prevent race conditions and double spending.

- **Explicit Dependency Injection**  
  Dependencies are constructed in one place (`app.Build`) and passed explicitly, following idiomatic Go practices.

---

## Project Structure

```text
cmd/
└── api/
    └── main.go           # application entry point

internal/
├── app/                  # dependency wiring and runtime setup
├── config/               # environment configuration
├── db/                   # database initialization (GORM)
├── domain/               # core business entities and rules
├── repository/           # data access layer
├── service/              # business workflows
├── handler/              # HTTP handlers
└── router/               # route registration
Tech Stack

Go

PostgreSQL

GORM

Gin (HTTP routing)

UUIDs for entity identity

Example Workflow: Wallet Transfer

Client submits a transfer request with an idempotency key.

Source wallet is locked at the database level.

Wallet balance is derived from ledger entries.

A transaction record is created.

Debit and credit ledger entries are written atomically.

Transaction is marked as completed.

If any step fails, the entire operation is rolled back.

Running the Project
Environment Variables
DB_URL=postgres://user:password@localhost:5432/fintech?sslmode=disable
PORT=8080

Start the Server
go run cmd/api/main.go

Design Philosophy

Prefer correctness over convenience

Make state changes explicit

Avoid hidden side effects

Keep business logic framework-agnostic

This codebase is structured to remain maintainable as complexity grows, and to reflect how financial systems are built in practice.

Status

This project is intended for learning, experimentation, and showcasing backend system design. It is not production-ready without further hardening, testing, and security review.


---

This README signals **engineering maturity**.  
Anyone who reads it understands immediately that you know how backend systems—and money—actually work.

Next logical upgrade:  
**add tests around the transfer service using a transactional test database**. That’s the line between “good project” and “serious engineer.”
