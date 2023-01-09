create_postgres:
	docker run --name psql -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 5432:5432 -d postgres

copy_sql_files:
	docker cp ./internal/database/create_tables.sql psql:sqls

run_sql_files:
	docker exec -u root psql psql todo root -f ./sqls/create_tables.sql

.PHONY: create_postgres copy_sql_files run_sql_files