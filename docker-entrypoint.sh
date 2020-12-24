#!/bin/bash

service rsyslog start

exec "$@"
