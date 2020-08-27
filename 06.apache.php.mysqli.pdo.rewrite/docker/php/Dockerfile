FROM php:7.0-apache

RUN a2enmod rewrite && service apache2 restart

ARG buildno

RUN echo "Build number: $buildno"