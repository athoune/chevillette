Chevillette
===========

> Tirez la bobinette et la chevillette cherra

You have two websites, one authenticated, the other is a preprod, almost public.

If your IP and user agent was seen in A server, you are whitelisted is site B.

It's a basic anti bot tools, not a real security tool, like basic auth or oauth2.
This tool doesn't handle very well coworking spaces or any public access point.

Web server
----------

Nginx with [http_auth_request](https://nginx.org/en/docs/http/ngx_http_auth_request_module.html).

Traefik with [ForwardAuth](https://doc.traefik.io/traefik/middlewares/http/forwardauth/).


Demo
----

Build your own `chevillette`

    make build-with-docker


### Fluentd demo

Go to demo folder

    cd demo-fluentd

### Loki demo

Launch services

    docker compose up -d

### Client

Test protected website B

    docker compose run client curl nginx-b

It should be 403

Unlock with website A

    docker compose run client curl nginx-a

Test website B again

    docker compose run client curl nginx-b

It should be 200

