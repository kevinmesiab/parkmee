# Parkmee.com

## _The AirBnb of parking - Rent your driveway_

www.parkmee.com attempts to create a distributed network of home owners, land owners and parking garage passes to lease thier spaces when not in use.

## Development Stragy

Parkmee is developed in Go, with an Angular front end.

Parkmee is distributed as a Docker file and simply cloning this repo.

## Local Development

Clone this repository.

1. Run `npm install`
2. Run `npm grunt`
3. Run `npm test`
4. Run `docker-compose up -db`

## Docker

### Dockerfile

Parkmee.com is built directly on the `golang:1.14` image
https://github.com/mesiablabs/www.parkmee.com/blob/develop/Dockerfile

### `docker-compose`

Clone this repository.
1. `npm install`
2. `docker-compose up`
4. Connect to http://localhost:8000

---
### License

License:
Copyright (C) 2020 Kevin Mesiab

    This program comes with ABSOLUTELY NO WARRANTY;

    This is free software, and you are welcome to redistribute it
    under certain conditions;
