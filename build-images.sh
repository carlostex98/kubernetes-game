#archivo para crear las imagenes de 
sudo docker login

#primera fase
cd grpc-client-kafka
sudo docker build -t  carlostex/grpc-client-kafka .
sudo docker push carlostex/grpc-client-kafka
cd ..

cd grpc-client-rabbit
sudo docker build -t  carlostex/grpc-client-rabbit .
sudo docker push carlostex/grpc-client-rabbit
cd ..

cd grpc-client-pubsub
sudo docker build -t  carlostex/grpc-client-pubsub .
sudo docker push carlostex/grpc-client-pubsub
cd ..

#segunda fase
cd kafka-client
sudo docker build -t  carlostex/kafka-client .
sudo docker push carlostex/kafka-client
cd ..

cd rabbitmq-client
sudo docker build -t  carlostex/rabbitmq-client .
sudo docker push carlostex/rabbitmq-client
cd ..

cd pubsub-client
sudo docker build -t  carlostex/pubsub-client .
sudo docker push carlostex/pubsub-client
cd ..

#tercera fase
cd kafka-worker
sudo docker build -t  carlostex/kafka-worker .
sudo docker push carlostex/kafka-worker
cd ..

cd rabbitMQ-worker
sudo docker build -t  carlostex/rabbitmq-worker .
sudo docker push carlostex/rabbitmq-worker
cd ..

cd pubsub-worker
sudo docker build -t  carlostex/pubsub-worker .
sudo docker push carlostex/pubsub-worker
cd ..




