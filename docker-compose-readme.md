Создать images из dockerfile
1. sudo docker build -f ./Dockerfile .
Тут уже докер композ
2. sudo docker-compose build
3. sudo docker-compose up registry 
4. sudo docker-compose -f ./docker-compose.yml build
5. sudo docker-compose -f ./docker-compose.yml push
6. sudo docker-compose -f ./docker-compose.yml up







удалить неиспользуемые images
sudo docker system prune

просмотреть images
sudo docker images

удалить images по идентификатору
sudo docker rmi 3a0f7b0a13ef



создать images из dockerfile
sudo docker build -f ./Dockerfile .
Successfully built 7818e805855c


запустить контейнер
sudo docker run 7818e805855c

удалить контейнеры
sudo docker stop name
sudo docker rm name 

запущенные контейнеры
sudo docker ps -a




sudo docker build .
sudo docker-compose -f ./docker-compose.yml up