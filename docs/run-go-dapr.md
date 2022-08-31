# Running Go-dapr

## Run Go-dapr from the CLI

The root folder of the repository contains [Docker Compose](https://docs.docker.com/compose/) files to run the solution locally. The `docker-compose.yml` file contains the definition of all the images needed to run go-dapr. The `docker-compose.override.yml` file contains the base configuration for all images of the previous file.

To start go-dapr from the CLI, run the following command from the root folder:

```terminal
docker-compose up
```

First Docker pulls the images. This can take some time to complete. Once the images are available, Docker will start the containers. You should now see the application logs in the terminal:















