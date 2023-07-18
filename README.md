# BuggyBuddy
BuggyBuddy is a Golang server that fails a percentage of responses with 500 Internal Server Error so you can test resiliency features.

Keyphrases:
* "_I need a microservice which I can make unstable for testing purposes_"
* "_Container that fails_"
* "_A server to configure how many requests it will fail_"

## Usage
Step in the root directory and
`docker build -t buggybuddy .`

Run in locally with
`docker run --publish 8000:8000 --rm buggybuddy`

### Methods

Working endpoint
`curl -v 'localhost:8000/'`

Failing endpoint
`curl -v 'localhost:8000/fail'`

Get the current failure percentage
`curl -v 'localhost:8000/getPercentage'`

Modify the failure percentage
`curl -v 'localhost:8000/setPercentage?percentage=30'`

Initial failure rate is 70%

## TODO
* Create endpoints to make responses slower
* Use a file to store config
* Load config using CM
