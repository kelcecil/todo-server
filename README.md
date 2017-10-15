# todo-server

## Description

Todo server is a simple server that allows users to add, list, or delete todos. This server is intended to be used for interviews as a simple API for a candidate to integrate against during an hour-long (or so) pairing session. 

## Operations of the API

The API can add, list, or delete todos. 

### Adding a todo

The endpoint for adding a todo can be found at `/add/`. The application should POST include a JSON body containing a `key` and `description`. The `key` is chosen by the front end and can be any valid string. The `description` is also a free-form string.

```json
{
    "key": "1",
    "description": "Write better Go code..."
}
```

Below is an example curl to write data to the 

```
curl -X POST -H "Accept: application/json" -d '{"key": "1", "description":"Write better Go code..."}' localhost:8080/add/
```

The service will return a 200 with no response body on a successful add. A 500 will be returned in the event of a problem.

### Listing todos

The endpoint to listing the current todos is at `/list/`. This should be a `GET` request with no parameters. An example curl request follows: 

```
curl -X GET localhost:8080/list/
```

The response is a JSON object with a single key `items` containing an array of objects representing todos. An example might look like:

```json
{
    "items": [
        {
            "key": "1",
            "description": "Write better Go code."
        },
        {
            "key": "2",
            "description": "Determine if Deckard is actually a replicant."
        }
    ]
}
```

### Deleting todos

The endpoint for adding a todo can be found at `/delete/`. The application should POST include a JSON body containing a `key`. All other keys in the request will be ignored. The API will currently always return 200. An example request body might look like:

```json
{
    "key": "1"
}
```

An example curl request to delete a key would look like: 

```
curl -X POST -H "Accept: application/json" -d '{"key": "1"}' localhost:8080/delete/
```

## Objective

Implement a JavaScript front-end (using libraries of your choice) to allow the user to interact with the API. Don't concentrate on completing within the time frame. Concentrate on producing work that you're proud of.

## Potential Questions

* How can the API be adjusted to provide communicate changes to the front-end? Can we make changes to ensure the front-end is more correct? To provide a better experience for the end user?
* What considerations do you think about when writing UI for front-end users? 
* What steps would you take to ensure a high quality front end (if given sufficient time)? Add tests? Collect errors from the browser?