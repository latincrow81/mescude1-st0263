FROM python:3.9.7-alpine

WORKDIR /usr/src/app
ADD pyproject.toml .
ADD poetry.lock .

RUN apk add build-base
RUN apk update && apk add postgresql-dev gcc python3-dev musl-dev curl py3-setuptools

RUN curl -sSL https://install.python-poetry.org | POETRY_VERSION=1.1.15 python3 -
RUN export PATH="/root/.local/bin:$PATH"
RUN /root/.local/bin/poetry config virtualenvs.create false
RUN if [ "$ENV" != "prd" ] ; then /root/.local/bin/poetry install; else /root/.local/bin/poetry install --no-dev; fi

ADD src src
ADD wsgi.py .
# Uncomment below if using Alembic for database migrations
# ADD alembic.ini .
# ADD alembic alembic
ADD gunicorn.config.py .
ADD config config

COPY docker/docker-entrypoint.sh /
RUN chmod +rx /docker-entrypoint.sh

RUN adduser -D --uid 1001 patientpop \
    && chown -R patientpop:patientpop /usr/src/app

USER 1001

EXPOSE 5010

ENTRYPOINT ["/docker-entrypoint.sh"]
