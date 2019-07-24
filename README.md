# Wait For Kibana

This is a container to use that waits for your Kibana service to be healthy before stopping. It does this by hitting the /api/status endpoint on the passed host and waiting for it to return a 200 [docs on this](https://www.elastic.co/guide/en/kibana/current/access.html).

[Kibana](https://www.elastic.co/products/kibana)

[Kibana User Guide](https://www.elastic.co/guide/en/kibana/current/index.html)
