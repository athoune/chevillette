Chevillette
===========

> Tirez la bobinette et la chevillette cherra

You have two websites, one authenticated, the other is a preprod, almost public.

If your IP and user agent was seen in A server, you are whitelisted is site B.

It's a basic anti bot tools, not a real security tool, like basic auth or oauth2.
This tool doesn't handle very well coworking spaces or any public access point.

Demo
----

    cd demo

Launch services

    docker compose up -d

Test protected website B

    docker run client curl nginx-b

It should 403

Unlock with website A

    docker run client curl nginx-a

Test website B again

    docker run client curl nginx-b

It should b 200
