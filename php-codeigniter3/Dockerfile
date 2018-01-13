# Start with our base image (the foundation) - version 7.1.5
FROM php:7.1.5-apache

# Install all the system dependencies and enable PHP modules
RUN apt-get update && apt-get install -y \
      libicu-dev \
      libpq-dev \
      libmcrypt-dev \
      git \
      zip \
      unzip \
    && rm -r /var/lib/apt/lists/* \
    && docker-php-ext-install \
      intl \
      mbstring \
      mcrypt \
      pcntl \
      pdo_pgsql \
      pgsql \
      zip \
      opcache

# Install composer
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/bin/ --filename=composer

# Set our application folder as an environment variable
ENV APP_HOME /var/www/html

# Change uid and gid of apache to docker user uid/gid
RUN usermod -u 1000 www-data && groupmod -g 1000 www-data

# Enable apache module rewrite
RUN a2enmod rewrite

# Copy the composer dependencies only
COPY app/composer.json $APP_HOME/composer.json

# Install all PHP dependencies without the autoloader
RUN composer install --no-ansi --no-dev --no-interaction --no-autoloader

# Set the port to 8080
RUN sed -i -e "s/VirtualHost \*:80/VirtualHost *:8080/g" /etc/apache2/sites-enabled/000-default.conf
RUN sed -i -e "s/Listen 80/Listen 8080/g" /etc/apache2/ports.conf

# Copy the app
COPY app $APP_HOME/

# Install all PHP dependencies
RUN composer install --no-ansi --no-dev --no-interaction --optimize-autoloader \
            && rm -rf /root/.composer/cache

# Change ownership of our applications
RUN chown -R www-data:www-data $APP_HOME
