#!/bin/bash

case $1 in
  create)
    curl -X POST \
      -H "Content-Type: application/json" \
      -d '{"name": "teste-01", "net": "full", "adb": "5555", "vnc": "6080"}' \
      localhost:8000/container/create
    
    curl -X POST \
      -H "Content-Type: application/json" \
      -d '{"name": "teste-02", "net": "lte", "adb": "5557", "vnc": "6081"}' \
      localhost:8000/container/create

    curl -X POST \
      -H "Content-Type: application/json" \
      -d '{"name": "teste-03", "net": "umts", "adb": "5559", "vnc": "6082"}' \
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
      echo "pass the id of the device as the second argument"
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

  update)
    if [[ $# != 2 ]]; then
      echo "pass the id of the device as the second argument"
      exit 1
    fi

    curl -X UPDATE \
      -H "Content-Type: application/json" \
      -d '{"name": "teste-update", "net": "full", "adb": "5561", "vnc": "6089"}' \
      localhost:8000/container/update/$2
  ;;

  *)
    echo -n "Usage: testing.sh <action> [argument]
      action: create, list, delete, update
      argument: device ID (need for delete and update)"
esac

