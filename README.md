### Image Shifter

This is a simple script that will shift the image to the right by 10 pixel per frame.

Use the Makefile to run the script. I recommend using the `make run-docker` command to run the script safely in a docker container, but change the resource limitaions in the docker-compose.yml file.

Check the constants in the script to change the number of frames and the concurrency of the script.