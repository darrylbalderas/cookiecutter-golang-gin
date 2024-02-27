IMAGE_NAME := cookiecutter-golang-gin:0.0.1

build:
	@docker build -t ${IMAGE_NAME} .
run:
	@docker run -p 8000:8000 ${IMAGE_NAME}