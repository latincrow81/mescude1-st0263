#!/bin/sh

set -e

# Uncomment below if using Alembic for database migrations
# alembic upgrade head

gunicorn -c gunicorn.config.py --enable-stdio-inheritance --reload wsgi
