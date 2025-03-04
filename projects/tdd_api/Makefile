test: docker_build_test
	docker compose up -d
	docker compose exec -T http go test `go list ./... | grep e2e_tests`
	docker compose down

unit_test:
	go test `go list ./... | grep -v e2e_tests` -v

docker_build:
	docker build -t service .

docker_build_test:
	docker build -t service_test --target=test .

docker_run:
	docker run --rm --publish 8080:8080 service

create_migration:
	migrate create -ext sql -dir ./migrations $(name)

migrate_up:
	migrate -source file://migrations -database 'postgres://tester:testing@localhost:5432/music?sslmode=disable' up

migrate_down:
	migrate -source file://migrations -database 'postgres://tester:testing@localhost:5432/music?sslmode=disable' down