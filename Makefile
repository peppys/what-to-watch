build:
	docker build -t  gcr.io/personal-200804/site-api:latest .
run:
	make build
	docker run -t -i --rm -p 8080:8080  gcr.io/personal-200804/site-api:latest
# push:
# 	gcloud docker -- push gcr.io/personal-200804/site-api:latest
# deploy:
# 	kubectl set image deployments personal-site-deployment personal-site-api=gcr.io/personal-200804/site-api:latest
# 	kubectl rollout status deployments personal-site-deployment