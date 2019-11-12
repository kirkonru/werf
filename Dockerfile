# syntax=docker/dockerfile:1.1.1-experimental
FROM ubuntu:16.04
COPY --from=composer:1.8.6 /usr/bin/composer /usr/bin/composer
RUN --mount=type=ssh date
