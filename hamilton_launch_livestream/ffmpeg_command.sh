#!/bin/bash


ffmpeg -f video4linux2 -s 640x480 -input_format mjpeg -r 30 -i /dev/video0 http://localhost:8090/feed1.ffm

