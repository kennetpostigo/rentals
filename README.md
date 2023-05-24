## Rentals

This repo contains the source for a Rentals API.

![Rental](https://giphy.com/embed/xfqDfDC2K3I3e)

## Implementation

Welcome! the implementation is mainly based on 2 libraries:

1. Chi - Http Router
2. Gorm - DB ORM

Building on those, it mainly was around managing query params, potential points of 
errors, DB Query, Query Param Validation, Pagination and more. I've got to say, touching go again 
for the first time in a while is very refreshing. Go literally gets out of the way and 
isn't too noisy or complainy about how you do much of anything.

## A few notes

Hi! A few things to keep in mind:

- I haven't written Go in a few years, so if there is any moments
  of "This piece of code doesn't seem to follow go convention/style" that is why.

- I'm very familiar with unit testing, and blackbox testing API's
  (I've had to write these quite often at work). However, I did not write
  any tests because today, I happened to have had a pretty shit day, and didn't feel up to
  googling go api testing best practices and then reading documentation.(I currently write mostly Rust and Node/Deno server-side code)

- I know the two items above aren't ideal, and I apologize about that. I hope that doesn't have an outsized impact on this opportunity, and if it does, I understand.

## Instructions

To run the project run:

```sh
docker compose up
```

### Potential Docker Warning/Error

1. If you get an error about the platform, update the platform in `docker-compose.yml` to that of your OS, I had to
   specify the platform b/c I'm running this on a macOS M1, and docker was complaining.

2. After you run the command, please wait for the `app` service to get to a healthy state, it will fail and restart, until
   postgres is ready to take on connections.

![Thats all folks](https://giphy.com/embed/xUPOqo6E1XvWXwlCyQ)
