# MySQL on Hasura

- Create a secret to hold root password:
  ```bash
  $ hasura secrets update mysql.password <secure-password>
  # you can verify it is set using 
  $ hasura secrets list
  ```
This password will be taken in by MySQL as `MYSQL_ROOT_PASSWORD`

Note: if you need to modify `conf`, uncomment line 4 in `Dockerfile`

### Other environment variables:

* `MYSQL_DATABASE` - This variable is optional and allows you to specify the name of a database to be created on image startup.
* `MYSQL_USER`, `MYSQL_PASSWORD` - These variables are optional, used in conjunction to create a new user and to set that user's password. Both variables are required for a user to be created.


### Connection parameters:

- username: `root`
- password: value set as `MYSQL_ROOT_PASSWORD`
- hostname: `mysql.default` (check by executing `$ hasura microservice list`)
- port: `3306`
