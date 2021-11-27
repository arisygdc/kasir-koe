installpg:
	docker run -d --name kasir_pg -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123qweasd -e POSTGRES_DB=kasir postgres:12-alpine3.14

uninstallpg:
	docker container rm kasir_pg

startpg:
	docker start kasir_pg

stoppg:
	docker stop kasir_pg

execdb:
	docker exec -it kasir_pg psql -U postgres

migrateup:
	migrate -path database/migration/ -database "postgresql://postgres:123qweasd@localhost:5432/kasir?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migration/ -database "postgresql://postgres:123qweasd@localhost:5432/kasir?sslmode=disable" -verbose down
  
.PHONY: installpg uninstallpg startpg stoppg execdb migrateup migratedown
