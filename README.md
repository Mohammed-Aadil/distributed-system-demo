# Distributes system demo

A project with following services

1. File uploader service for media files (golang) [`storage`]:
   It have endpoint to upload and download doc file. It communicates to `doc_to_png` and `postgres` services. It also save attachment data in database.

2. doc, txt, pdf to png converter (python) [`doc_to_png`]:
   It is simple flask server to convert doc to images.

3. client (react) [`client`]:
   It is pure react base client with communicate to `storage` service. It allows you to upload and download doc.
4. Postgresql (docker-image) [`postgres`]:
   It is postgresql database pulled from official psql repo of docker hub.

## How to run

1. `cd distributed-system-demo`
2. `docker-compose up`

That's all you need to do.
