include ./.env
DBURL=postgres://$(DBUSER):$(DBPASS)@$(DBHOST):$(DBPORT)/$(DBNAME)?sslmode=disable
MIGRATIONPATH=db/migrations

migrate-createUp:
	migrate -database $(DBURL) -path $(MIGRATIONPATH) up

migrate-createDown:
	migrate -database $(DBURL) -path $(MIGRATIONPATH) down

migrate-status:
	migrate -database $(DBURL) -path $(MIGRATIONPATH) version

migrate-force:
	migrate -database $(DBURL) -path $(MIGRATIONPATH) force $(v)