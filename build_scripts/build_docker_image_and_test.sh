cd $BUILD_DIR/$APP_NAME
docker build -t ${DOCKER_IMAGE_ID} .
docker run --name ${DOCKER_CONTAINER_ID} ${DOCKER_IMAGE_ID} go test
