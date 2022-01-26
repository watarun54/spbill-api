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
	aws s3 cp server.zip s3://serverless-skill-manager/ &&\
	rm server && rm server.zip

deploy: build zip
	cd api &&\
	aws lambda update-function-code --function-name serverless-skill-manager --zip-file fileb://server.zip &&\
	rm server && rm server.zip
