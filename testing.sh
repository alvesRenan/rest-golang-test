#!/bin/bash

case $1 in
  create)
    curl -X POST \
      -H "Content-Type: application/json" \
      -d '{"name": "teste-01", "net": "full", "serial_port": "5554", "adb_port": "5555", "vnc_port": "6080"}' \
      localhost:8000/container/create
    
    curl -X POST \
      -H "Content-Type: application/json" \
      -d '{"name": "teste-02", "net": "lte", "serial_port": "5556", "adb_port": "5557", "vnc_port": "6081"}' \
      localhost:8000/container/create

    curl -X POST \
      -H "Content-Type: application/json" \
      -d '{"name": "teste-03", "net": "full", "serial_port": "5558", "adb_port": "5559", "vnc_port": "6082"}' \
      localhost:8000/container/create
  ;;
  
  create-scenario)
    curl -X POST \
      -H "Content-Type: application/json" \
      -d '{"name": "scenario-01", "state": "created"}' \
      localhost:8000/scenario/create
    curl -X POST \
      -H "Content-Type: application/json" \
      -d '{"name": "scenario-02", "state": "created"}' \
      localhost:8000/scenario/create
    curl -X POST \
      -H "Content-Type: application/json" \
      -d '{"name": "scenario-03", "state": "created"}' \
      localhost:8000/scenario/create
  ;;

  list)
    curl -X GET \
      localhost:8000/container/list
  ;;

  list-scenario)
    curl -X GET \
      localhost:8000/scneario/list
  ;;

  delete)
    if [[ $# != 2 ]]; then
      echo "pass the name of the container as the second argument"
      exit 1
    fi

    curl -X DELETE \
      -H "Content-Type: application/json" \
      localhost:8000/container/delete/$2
  ;;

  delete-scenario)
    if [[ $# != 2 ]]; then
      echo "pass the name of the scenario as the second argument"
      exit 1
    fi

    curl -X DELETE \
      -H "Content-Type: application/json" \
      localhost:8000/scenario/delete/$2
  ;;

  *)
    echo -n "Usage: testing.sh <action> [argument]
      action: create, list, delete
      argument: container name (needed for delete)"
esac

