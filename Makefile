.PHONY: help

# Examples of environment variables
# DROPS_PRESCRIPTION=2s->http://localhost:8099/ping;4s->http://localhost:8099/pong

run-pinger:
    PORT=8099 ./build/dropper

run-dropper:
    PORT=8097 DROPS_PRESCRIPTION="2s->http://localhost:8099/ping;4s->http://localhost:8099/pong" ./build/dropper
