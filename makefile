buildlatest:
	docker build -t test:latest .

runlatest:
	docker run -p 9000:9000 test:latest