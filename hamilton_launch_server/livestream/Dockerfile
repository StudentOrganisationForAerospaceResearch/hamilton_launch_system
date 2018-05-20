FROM ubuntu:16.04 AS base

WORKDIR /hamilton_launch_livestream/

# From https://trac.ffmpeg.org/wiki/CompilationGuide/Ubuntu
RUN apt-get update -qq && apt-get -y install \
        autoconf \
        automake \
        build-essential \
        cmake \
        git \
        libass-dev \
        libfreetype6-dev \
        libsdl2-dev \
        libtheora-dev \
        libtool \
        libva-dev \
        libvdpau-dev \
        libvorbis-dev \
        libxcb1-dev \
        libxcb-shm0-dev \
        libxcb-xfixes0-dev \
        mercurial \
        pkg-config \
        texinfo \
        wget \
        zlib1g-dev \
        yasm \
        v4l-utils \
        ffmpeg

#RUN cd FFmpeg && \
        #git checkout 2ca65fc && \
        #./configure && \
        #make -j8 && \
        #make install
