include .env
export $(shell sed 's/=.*//' .env)

epoch=$(shell date +%s)

.PHONY: schema
schema:
	@touch $(DATABASE_MIGRATIONS_PATH)"/"$(epoch)"_[RENAME].up.sql"
	@touch $(DATABASE_MIGRATIONS_PATH)"/"$(epoch)"_[RENAME].down.sql"
	@echo ">> Created schema in "$(DATABASE_MIGRATIONS_PATH)
	@echo "\t- "$(epoch)"_[RENAME].up.sql"
	@echo "\t- "$(epoch)"_[RENAME].down.sql"
	@echo ">> Please rename the generated files"

.PHONY: migrate-up
migrate-up:
	@migrate -source file://${DATABASE_MIGRATIONS_PATH} -database ${DATABASE_URL} up

.PHONY: migrate-down
migrate-down:
	@migrate -source file://${DATABASE_MIGRATIONS_PATH} -database ${DATABASE_URL} down 1