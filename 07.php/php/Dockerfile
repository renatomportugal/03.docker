FROM php:7.0-apache

# RUN apk update 
# \
#     && a2enmod rewrite \
    # && systemctl restart apache2

# FROM httpd:2.4.33-alpine
# RUN apk update; \
#     apk upgrade;

RUN a2enmod rewrite && service apache2 restart

ARG buildno

RUN echo "Build number: $buildno"