#!/bin/bash

# Add Widget
data=
curl -i \
-X POST \
-d '{"Id": "X", "Row": "1", "Col": "1", "Sizex": "1", "Sizey": "1", "Content": "Example Dashpersonal"}' \
-H "Content-type: application/json" "http://localhost:9000/"

# Update widget
data=
curl -i \
-X POST \
-d '{"Id": "2", "Content": "HEYY"}' \
-H "Content-type: application/json" "http://localhost:9000/"
