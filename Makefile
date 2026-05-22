include .env
export


service-run:
	go run ./main.go
	cmd1
	cmd2
	cmd3

some-target:
	export env_var=hello && \
	printenv env_var

migrate-up:
	migrate -path migrations -database ${CONN_STRING} up

migrate-down:
	migrate -path migrations -database ${CONN_STRING} down

migrate-force:
	migrate -path migrations -database ${CONN_STRING} force 0
