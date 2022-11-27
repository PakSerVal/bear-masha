PACKAGE=github.com/PakSerVal/bear-masha/cmd/bear

run-local-server:
	docker compose up -d && \
	cd server && \
	goose -dir migrations postgres "host=localhost port=5432 user=postgres password=postgres sslmode=disable" up && \
	go run ${PACKAGE} -secretKey="some secret"

run-client:
	cd client && npm install && npm start
