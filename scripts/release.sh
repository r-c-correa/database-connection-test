VERSION=0.0.2
CONTAINER_NAMESPACE=ricardocorrea
SOFTWARE=database-connection-test

docker build -f ./scripts/Dockerfile -t $CONTAINER_NAMESPACE/$SOFTWARE:$VERSION .

docker push $CONTAINER_NAMESPACE/$SOFTWARE:$VERSION

echo "Finished. Press a key to exit..."
read -s -n 1 key