build:
	docker build -t restgo:latest .

run:
	docker run -p 8080:8080 restgo:latest