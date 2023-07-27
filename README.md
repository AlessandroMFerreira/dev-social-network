# DEV-SOCIAL-NETWORK

This project was built to test my golang skills. It's a social network for developers.

**Golang version:**

    1.19

# How to run this application
### Run local using docker containers 
In root directory run:
> $ docker-compose up

The command above will start an mysql data base and the application listening on port **5000**

### Run local
In this case you will have to provide an mysql database (see the docker-compose for data base credentials. You can alsol comment the part responsable to build the api container on docker-compose file and run it).
> $ cd api/

> $ go run main.go
