# Wait For Kibana

Find me on [Docker Hub](https://hub.docker.com/r/mmcken3/wait-for-kibana)

This is a container to use that waits for your Kibana service to be healthy before stopping. It does this by hitting the /api/status endpoint on the passed host and waiting for it to return a 200 ([docs on this](https://www.elastic.co/guide/en/kibana/current/access.html)).

[Kibana](https://www.elastic.co/products/kibana)

[Kibana User Guide](https://www.elastic.co/guide/en/kibana/current/index.html)

## Run With Container

Running the container with no host flag uses the default host of [http://localhost:5601](http://localhost:5601).

```bash
docker run --network host mmcken3/wait-for-kibana
```

You can pass in host flags like below:

```bash
docker run --network host mmcken3/wait-for-kibana -host http://yourhost
```

## Run Outside Container

This is how you can run this locally outside of a container.

```bash
go build -o waitforkibana main.go
./waitforkibana -host http://yourhost
```

## Push To Docker Hub

[Image In Docker Hub](https://hub.docker.com/r/mmcken3/wait-for-kibana)

Get the current github short tag using this command:

```bash
git describe --abbrev=0 --tags
```

Build, Tag, and push this Docker image to Docker Hub using that tag.

```bash
docker build -t mmcken3/wait-for-kibana .
docker tag mmcken3/wait-for-kibana:latest mmcken3/wait-for-kibana:[github-tag]
docker push mmcken3/wait-for-kibana:latest
docker push mmcken3/wait-for-kibana:[github-tag]
```
