run:
	docker build -t personal-image .
	docker run -t --rm -p 80:8080 personal-image
push:
	docker tag personal-image xpeppy/personal-site-api
	docker push xpeppy/personal-site-api