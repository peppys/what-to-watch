run:
	docker build -t personal-image .
	docker run -t -i --rm -p 8080:80 personal-image
push:
	docker tag personal-image xpeppy/personal-site-api
	docker push xpeppy/personal-site-api