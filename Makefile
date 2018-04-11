run:
	docker build -t personal-image .
	docker run -t -i --rm -p 80:80 -p 443:443 personal-image
push:
	docker tag personal-image xpeppy/personal-site-api
	docker push xpeppy/personal-site-api