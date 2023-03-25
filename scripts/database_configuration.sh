#!/bin/bash

docker exec -i database psql -U user -w -c "CREATE DATABASE error_tracker"
