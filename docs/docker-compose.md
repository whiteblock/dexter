---
id: docker-compose
title: Getting Started With docker-compose
---

# Getting Everything Running At Once

To make things easier for everyone, a `docker-compose.yml` has been provided that gets everything running
in one easy step.

```sh
docker-compose up 
```

This gives you the following:

* PostgreSQL running on port 5432
* dexter-data running on port 50051
* grpcwebproxy running on port 8081 as a proxy for dexter-data
* dexter running on port 50052
* grpcwebproxy running on port 8082 as a proxy for dexter

The first time will take a while to download all the dependencies, but subsequent startup times should be relatively quick.
