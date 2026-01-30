MIGRATIONS_FOLDER = db/migrations
DB_URL = root:\#MBAV@584@tcp(127.0.0.1:3306)/auth_dev

# Create a new migration -> make migrate-create name="create_user_table"
migrate-create:
	goose -dir $(MIGRATIONS_FOLDER) create $(name) sql

migrate-up:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" up

migrate-down:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" down

migrate-reset:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" reset

migrate-status:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" status

migrate-redo:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" redo

migrate-to:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" up-to ${version}

migrate-down-to:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" down-to ${version}

migrate-force:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" force ${version}

migrate-help:
	goose -dir $(MIGRATIONS_FOLDER) help