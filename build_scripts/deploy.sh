echo "Deploying ${DOCKER_IMAGE_ID} for $1"
python deploy.py install dev01 --chart=heads_or_tails --verbose
