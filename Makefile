up:
	docker compose up -d --build

down:
	docker compose down

setup:
	@./scripts/setup.sh

sh:
	docker exec -it app sh

db:
	docker exec -it database psql -U user -w
