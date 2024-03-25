start:
	go build -o fishing cmd/web/*.go 
	./fishing

docker_db_build:
	docker build ./db/ -t fishing-postgres
docker_db_start:
	docker run --name fishing-postgres -p 5432:5432 -v "./db/data/postgres/:/var/lib/postgresql/data:Z" -d fishing-postgres
docker_db_stop:
	docker stop fishing-postgres
	docker rm fishing-postgres

