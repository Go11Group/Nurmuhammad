DB_URL='postgres://postgres:pass@localhost:5432/new?sslmode=disable'


create_migrate:
	@echo "Enter file name: "; \
	read filename; \
	migrate create -ext sql -dir migrations -seq $$filename
up:
	migrate -database ${DB_URL} -path migrations up
down:
	migrate -database ${DB_URL} -path migrations down
fix:
	@echo "Enter version number: "; \
	read num; \
	migrate -database ${DB_URL} -path migrations force $$num
goto:
	migrate -database ${DB_URL} -path migrations goto 1
mod:
	@echo "Enter module name: "; \
	read modulename; \
	go mod init $$modulename
	go mod tidy
