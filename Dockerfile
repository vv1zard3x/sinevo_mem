FROM python:3.12.4-slim

SHELL ["/bin/bash", "-c"]

EXPOSE 8000

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONBUFFERED 1
ENV POETRY_VIRTUALENVS_CREATE=false

RUN pip install --upgrade pip
RUN pip install poetry
RUN apt update && apt -qy install gcc libjpeg-dev 

RUN useradd -rms /bin/bash user && chmod 777 /opt /run

WORKDIR /code
RUN mkdir /code/static && mkdir /code/media
COPY --chown=user:user . .
RUN chown -R user:user /code && chmod -R 755 /code

RUN poetry config virtualenvs.create false \
    && poetry install --no-interaction --no-ansi --no-root

USER user

CMD ["poetry", "run", "gunicorn", "-b", "0.0.0.0:8000", "app.wsgi:application"]
