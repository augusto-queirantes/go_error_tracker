build:
	docker compose up -d --build

up:
	docker compose up -d --build

down:
	docker compose down

setup:
	@./scripts/setup.sh

config_database:
	@./scripts/database_configuration.sh

sh:
	docker exec -it app sh

db:
	docker exec -it database psql -U user -w -d error_tracker
