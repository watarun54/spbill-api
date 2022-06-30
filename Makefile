CONTAINER := echo

.PHONY: deps clean build

deps:
	go get -u ./...
clean: 
	rm -rf ./api/server
build:
	cd api && GOOS=linux GOARCH=amd64 go build
zip:
	cd api && zip server.zip server
upload-s3: build zip
	cd api &&\
	aws s3 cp server.zip s3://spbill-api/ &&\
	rm server && rm server.zip
deploy: build zip
	cd api &&\
	aws lambda update-function-code --function-name spbill-api --zip-file fileb://server.zip &&\
	rm server && rm server.zip
ps:
	docker-compose ps
logs:
	docker-compose logs -f --tail=500
up:
	docker-compose up -d
down:
	docker-compose down
down-rm:
	docker-compose down --rmi all --volumes
restart:
	docker-compose restart $(CONTAINER)
exec:
	docker-compose exec $(CONTAINER) sh
