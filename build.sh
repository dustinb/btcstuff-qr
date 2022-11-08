docker build -t oldbute/btcstuff-qr:latest .
docker push oldbute/btcstuff-qr:latest

kubectl apply -n btcstuff -f kube
kubectl rollout restart deployment/btcstuff-qr
