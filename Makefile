dockerbuild:
	docker build --build-arg RELEASE=1 -t alice:latest .
