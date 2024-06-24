DB_URL='postgres://postgres:pass@localhost:5432/new?sslmode=disable'
DB_URL_2='postgres://postgres:pass@localhost:5432/lesson43?sslmode=disable'




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


create_migrate_2:
	@echo "Enter file name: "; \
	read filename; \
	migrate create -ext sql -dir migration -seq $$filename
up_2:
	migrate -database ${DB_URL_2} -path migration up
down_2:
	migrate -database ${DB_URL_2} -path migration down
fix_2:
	@echo "Enter version number: "; \
	read num; \
	migrate -database ${DB_URL_2} -path migration force $$num
goto_2:
	migrate -database ${DB_URL_2} -path migration goto 1
