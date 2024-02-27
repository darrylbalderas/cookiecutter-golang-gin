IMAGE_NAME := {{ cookiecutter.project_slug }}:{{ cookiecutter.version }}

build:
	@docker build -t ${IMAGE_NAME} .
run:
	@docker run -p {{ cookiecutter.project_port }}:{{ cookiecutter.project_port }} ${IMAGE_NAME}