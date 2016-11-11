build:
	go build -o pgdock
run: build
	docker run -p 5432:5432 -d --name pgdock_db gwhn/postgres:dev
	./pgdock &
	docker logs -f pgdock_db
migrate:
	migrate -url "postgres://test_user:1234@localhost:5432/test_db?sslmode=disable" \
		-path ./db/mig up
stop:
	killall pgdock
	docker stop pgdock_db
	docker rm pgdock_db
