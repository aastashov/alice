dockerbuild:
	docker build --build-arg RELEASE=0.0.1 --platform linux/amd64 -t aastashov/alice:latest .

dockerpush:
	docker push aastashov/alice:latest

release: dockerbuild dockerpush
