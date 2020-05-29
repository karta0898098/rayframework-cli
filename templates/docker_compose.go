package templates

const DockerCompose = `
version: '3.7'

services: 
    database:
        image: mysql:8.0.19
        command: --default-authentication-plugin=mysql_native_password
        restart: always
        environment:
            MYSQL_ROOT_PASSWORD: %s
        ports: 
            - "8081:3306"
        volumes: 
            - ./sql:/docker-entrypoint-initdb.d
            - ./data:/var/lib/mysql
    redis:
        image: "redis:alpine"
        
        command: redis-server --requirepass %s
        
        ports:
            - "8082:6379"

    backend-app:
        build: .
        depends_on: 
            - database
            - redis
        ports: 
            - "8080:8080"
        environment: 
            TZ: Asia/Taipei
            REDIS_ADDRESS: redis:6379
            REDIS_PASSWORD: %s
            DB_CONNECTION_STRING: root:%s@(database:3306)/%s?charset=utf8&parseTime=True&loc=Local
`

