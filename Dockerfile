FROM python:3.6.4

RUN pip install \
    flask==0.12.2 \
    pyflakes==1.6.0

WORKDIR /hamilton_launch_system
