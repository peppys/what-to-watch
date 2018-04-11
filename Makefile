build:
	docker build -t  gcr.io/personal-200804/site-api:v$(v) .
run:
	make build
	docker run -t -i --rm -p 8080:8080  gcr.io/personal-200804/site-api:v$(v)
push:
	gcloud docker -- push gcr.io/personal-200804/site-api:v$(v)
deploy:
	kubectl apply -f kubernetes.yaml
	kubectl rollout status deployments personal-site-deployment