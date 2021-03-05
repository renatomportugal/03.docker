# Dockerfile for Dashboard Profile of WSO2 API Manager Analytics #

This section defines the step-by-step instructions to build [Alpine](https://hub.docker.com/_/alpine/) Linux based Docker image for Dashboard profile of
WSO2 API Manager Analytics 3.2.0.

## Prerequisites

* [Docker](https://www.docker.com/get-docker) v17.09.0 or above
* [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) client

## How to build an image and run

##### 1. Checkout this repository into your local machine using the following Git client command.

```
git clone https://github.com/wso2/docker-apim.git
```

> The local copy of the `dockerfile/alpine/apim-analytics/dasboard` directory will be referred to as `ANALYTICS_DOCKERFILE_HOME` from this point onwards.

##### 2. Build the Docker image.

- Navigate to `<ANALYTICS_DOCKERFILE_HOME>` directory. <br>
  Execute `docker build` command as shown below.
    + `docker build -t wso2am-analytics-dashboard:3.2.0-alpine .`
    
> By default, the Docker image will prepackage the General Availability (GA) release version of the relevant WSO2 product.
    
##### 3. Running Docker images specific to each profile.

- `docker run -p 9643:9643 wso2am-analytics-dashboard:3.2.0-alpine`

> Here, only port 9643 has been mapped to a Docker host port.
You may map other container service ports, which have been exposed to Docker host ports, as desired.

##### 4. Accessing the Dashboard portal.

- For dashboard,
    + `https:<DOCKER_HOST>:9643/analytics-dashboard`
    
> In here, <DOCKER_HOST> refers to hostname or IP of the host machine on top of which containers are spawned.

## How to update configurations

Configurations would lie on the Docker host machine and they can be volume mounted to the container. <br>
As an example, steps required to change the port offset using `deployment.yaml` is as follows:

##### 1. Stop the API Manager Analytics container if it's already running.

In WSO2 API Manager Analytics version 3.2.0 product distribution, `deployment.yaml` configuration file <br>
can be found at `<DISTRIBUTION_HOME>/conf/dashboard`. Copy the file to some suitable location of the host machine, <br>
referred to as `<SOURCE_CONFIGS>/deployment.yaml` and change the `offset` value under `ports` to 1.

##### 2. Grant read permission to `other` users for `<SOURCE_CONFIGS>/deployment.yaml`.

```
chmod o+r <SOURCE_CONFIGS>/deployment.yaml
```

##### 3. Run the image by mounting the file to container as follows:

```
docker run 
-p 9641:9641
--volume <SOURCE_CONFIGS>/deployment.yaml:<TARGET_CONFIGS>/deployment.yaml
wso2am-analytics-dashboard:3.2.0-alpine
```

> In here, <TARGET_CONFIGS> refers to /home/wso2carbon/wso2am-analytics-3.2.0/conf/dashboard folder of the container.

## Docker command usage references

* [Docker build command reference](https://docs.docker.com/engine/reference/commandline/build/)
* [Docker run command reference](https://docs.docker.com/engine/reference/run/)
* [Dockerfile reference](https://docs.docker.com/engine/reference/builder/)
