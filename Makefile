.PHONY: up down rebuild logs test lint

up:
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
	fi
	docker compose up -d

down:
	docker compose down

fresh: rebuild migrate

rebuild:
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
	fi
	docker compose down --remove-orphans
	docker compose build --no-cache
	docker compose up -d --build -V

migrate:
	docker compose exec sprezz-web go run . migrate up

logs:
	docker compose logs -f

test:
	go test -v -race -cover -count=1 -failfast ./...

lint:
	golangci-lint run -v
