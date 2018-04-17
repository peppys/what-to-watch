build:
	docker build -t gcr.io/personal-200804/site-$(env):v$(v) web
	docker build -t gcr.io/personal-200804/api-$(env):v$(v) api
run:
	make build v=$(v) env=dev
	docker run -t -i --rm -p 8080:8080 -e VERSION=$(v) -e APP_ENV=dev gcr.io/personal-200804/site-api-dev:v$(v)
dev:
	make build v=$(v) env=dev
	make push v=$(v) env=dev
	kubectl config set current-context minikube
	# kubectl create secret docker-registry gcr \
    #     --docker-server=https://gcr.io \
    #     --docker-username=oauth2accesstoken \
    #     --docker-password=$(gcloud auth print-access-token) \
    #     --docker-email=peppysisay@gmail.com
	# kubectl patch serviceaccount default \
    #     -p '{"imagePullSecrets": [{"name": "gcr"}]}'
	kubectl apply -f local-kubernetes.yaml
push:
	gcloud docker -- push gcr.io/personal-200804/site-$(env):v$(v)
	gcloud docker -- push gcr.io/personal-200804/api-$(env):v$(v)
deploy:
	make build v=$(v) env=prod
	make push v=$(v) env=prod
	gcloud config set project personal-200804
	gcloud container clusters get-credentials cluster-1
	kubectl apply -f production-kubernetes.yaml
	kubectl rollout status deployments personal-site-deployment