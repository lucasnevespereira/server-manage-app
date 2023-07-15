.PHONY: run-front run-api build-api stop-api test-api

run-front:
	cd server-manage-front && pnpm i && pnpm run dev

test-front:
	cd server-manage-front && pnpm run test

build-api:
	cd server-manage-api && docker-compose up --build

run-api:
	cd server-manage-api && docker-compose up

stop-api:
	cd server-manage-api && docker-compose down

test-api:
	cd server-manage-api && go test -v ./...

test-coverage:
	cd server-manage-api && make test-coverage
