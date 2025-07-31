PORT ?= 8000
ARGS_TEST ?=
ARGS_SERVE ?=

.PHONY: install
install:
	uv sync

.PHONY: reinstall
reinstall: clean install

.PHONY: update
update:
	uv sync --upgrade

.PHONY: test
test:
	uv run ${ARGS_TEST} -m pytest

.PHONY: coverage
coverage:
	uv run -m coverage erase
	uv run -m coverage run -m pytest tests
	uv run -m coverage combine
	uv run -m coverage report
	uv run -m coverage html

.PHONY: lint
lint:
	uvx ruff check
	uvx ruff format

.PHONY: typing
typing:
	uvx ty check --python .venv src

.PHONY: check-all
check-all: lint typing coverage

.PHONY: serve
serve:
	uv run ${ARGS_SERVE} -m fastapi dev src/hello_svc/asgi.py --port ${PORT}

.PHONY: test-request
test-request:
	uvx --from httpie http http://127.0.0.1:${PORT}/

.PHONY: browser
browser:
	uv run -m webbrowser -t http://127.0.0.1:${PORT}/

.PHONY: clean
clean:
	rm -rf .venv .pytest_cache .ruff_cache .coverage htmlcov
	find . -type d -name "__pycache__" -exec rm -r {} +
