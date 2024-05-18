java: mongodb
	java -cp .:lib/* Queries.java

python: mongodb
	python3 queries.py

nodejs: mongodb
	node queries.js

go: mongodb
	go run queries.go

mongodb: clean
	docker run -d --name mongodb -p 127.0.0.1:27017:27017 mongo

clean:
	docker rm -f mongodb

.PHONY: java python nodejs go mongodb clean
