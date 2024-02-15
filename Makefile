local-environment:
	docker-compose -f docker-compose-run-local.yaml down
	docker-compose -f docker-compose-run-local.yaml stop
	docker-compose -f docker-compose-run-local.yaml build --force-rm
	docker-compose -f docker-compose-run-local.yaml up
