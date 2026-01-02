all:
	@echo usage:
	@echo   make deploy-container
	@echo   make deploy-container-kind

deploy-container:
	nix build .#external-dns-bunny-webhook-docker.stream-layered && ./result | doas ctr -n k8s.io image import -

deploy-container-kind:
	nix build .#packages.aarch64-linux.external-dns-bunny-webhook-docker.stream-layered && \
		sudo ssh rosetta-builder $$(readlink result) | docker load
                
