# Distributes system demo

A project with following services. This project uses all services as git sub-modules if you clone repo and wanted to inspect code of below services, clone sub-modules too.

1. File uploader service for media files (golang) [`storage`]:
   It have endpoint to upload and download doc file. It communicates to `doc_to_png` and `postgres` services. It also save attachment data in database.

2. doc, txt, pdf to png converter (python) [`doc_to_png`]:
   It is simple flask server to convert doc to images.

3. client (react) [`client`]:
   It is pure react base client with communicate to `storage` service. It allows you to upload and download doc.
4. Postgresql (docker-image) [`postgres`]:
   It is postgresql database pulled from official psql repo of docker hub.

## How to run

1. `git clone git@github.com:Mohammed-Aadil/distributed-system-demo.git --recurse-submodules`
2. `cd distributed-system-demo`
3. `docker-compose up` or `docker stack deploy -c docker-compose.yml ds`
4. To close all services `docker-compose down` or `docker stack rm ds`

That's all you need to do.
