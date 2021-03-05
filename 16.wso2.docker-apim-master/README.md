# Docker and Docker Compose Resources for WSO2 API Management

This repository contains following Docker resources:

- Alpine, CentOS and Ubuntu based Docker resources for WSO2 API Manager, API Manager Analytics Dashboard and Worker and
Identity Server as Key Manager profiles

- Docker Compose resources for the most common WSO2 API Management deployment patterns

Docker resources for WSO2 API Manager, API Manager Analytics and WSO2 Identity Server as Key Manager
help you build generic Docker images for deploying the corresponding product servers in containerized environments.

Each Docker image includes the Java Development Kit, the relevant product distribution and a collection of utility libraries.
Configurations and non-configuration resources (e.g. binaries such as, third-party libraries, Carbon extensions,
Carbon Applications and security related artifacts such as, Java Keystore files) are designed to be provided via
volume mounts to the containers spawned.

Docker Compose files have been created according to the most common API Management deployment patterns available for allowing users
to quickly evaluate product features along side their co-operate API Management requirements. The Compose files make use of per profile
Docker images of WSO2 API Manager, API Manager Analytics and WSO2 Identity Server as Key Manager, as well as MySQL.

**Change log** from previous v3.1.0.3 release: [View Here](CHANGELOG.md)
