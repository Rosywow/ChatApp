docker stop chatApp_frontend > /dev/null 2> /dev/null || :
docker container rm chatApp_frontend > /dev/nul 2> /dev/null || :

docker run -d --restart=always \
        -p 8088:80  \
        -p  443:433 \
 --name chatApp_frontend \
 -v deploy:/var/deploy:z \
 -v /home/rosy/myweb/config/nginx.conf:/etc/nginx/nginx.conf \
 -v /home/rosy/myweb/config/conf.d:/etc/nginx/conf.d \
 -v /home/rosy/myweb/logs:/var/log/nginx \
 --network chatnet \
 nginx:first

 