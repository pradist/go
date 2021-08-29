deploy-all:
	kubectl apply -f ./deploy/rockme-dc.yaml
	kubectl apply -f ./deploy/rockme-svc.yaml

delete-all:
	kubectl delete deploy rockme
	kubectl delete svc rockme-svc
