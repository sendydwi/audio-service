## Audio Service
# Dependency
Run on docker
- docker
    - how to install : https://docs.docker.com/get-started/get-docker/

Run on local dependency
- go
    - how to install : https://go.dev/doc/install
- ffmpeg
    - how to install : https://www.ffmpeg.org/download.html
- postgres
    - how to install : https://www.postgresql.org/download/

# How To Run using docker
- open folder using terminal
- run docker command `docker compose up -d` 
- to stop run docker command `docker compose down`

# How to Run without docker
- install local all dependency
- setup DB as specified in .env file
- run command go run ./cmd/main.go 

## Reasoning for Stack
- why go ?
    - go have great performance and have new perspective on simplicity and concurrency. Main reason why I choose this language is because i want to learn it. 
- why ffmpeg ?
    - ffmpeg are great open source tools to process audio, it support variety of encoding and it have high security measurement. this was one of the best option on audio processing
- why use godub ?
    - godub the only reliable open library i find. since it's use ffmped behind it make it reliable and support extensive use cases.
- why use postgres ?
    - no particular reason. I just have familiarity with it. it's relatively easy to setup and work for a lot of use cases. therefore it's my go to for simple application that need database
- why use automigration ?
    - this to simplified process on spawning infrastructure. for production cases it would be better to use migration and create separate repo for it to be easier to manage.
- why not use in memmory database ?
    - at first i want to use in memory database. but I may harder to test since the data will removed once the system shut down. therefore i choose to use to separate the database to it's onw instance
- why local storage ?
    - split between the audio conversion service and storage service to handle it would be better in real world application like using s3 or cloud storage for the storage service. but in this case the instance will have redundant api and may look more like overengineered so I think simplified it to single service will looks better
- why docker compose ?
    - docker compose great tools to manage multiple container at once and help to make collaboration easier since every engineers now could just use the docker compose to run the service. but int not meant to be use in production since it may make scaling complicated. in production i would use IaC like Terraform or Pulumi to setup the infrastructure and manage the container.


## Corners cutted:
# Design
- store data to local storage
- not using proper config / enviroment variable
- not use middle ware to better logging and auth
- not using migration
- not make external storage instances

# Code 
- unit test only on certain part and not test every edge cases
- not verifiying user id and phrase id
- not limit file size
- error handling not return proper error code

# Infra
- copy entire code instead of only it's executable
- not implement tracing and monitoring
- not using IaC
- not setup proper infra structure