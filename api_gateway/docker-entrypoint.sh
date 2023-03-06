#!/bin/sh

set -e

gunicorn --enable-stdio-inheritance --reload wsgi:app
