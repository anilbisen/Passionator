# passionator for Golang #
A simple stateless REST service in Golang. Build in Drone.

[![Build Status](http://164.177.187.96:18000/api/badges/okoeth/passionator-golang/status.svg)](http://164.177.187.96:18000/okoeth/passionator-golang)

## Set-up ##

### Installation of basics ###
Install golang from [here](http://golang.org). And set environment variable. In my case:
```
GOPATH=~/go
```

As editor I recommend `vi(1)` (just kidding) or for UI afficionados [VS Code](https://code.visualstudio.com/). Make sure you have the Go Extension installed.

Of course we also assume that [Git](https://git-scm.com/) is installed and that in case of Windows you use the Git Bash. Stay away from cmd.exe.

### Run locally ###
The service can be started locally (assuming a local Mongo is running, see Docker info below for running it):
```
$ go run *.go
```

Just a note: You were not bothered with downloading any libs or packages. This is because we use /vendoring/ which has been set-up using [Glide](http://glide.sh/). Consider installing Glide yourself.

Access REST API:
```
$ curl http://localhost:8000/v1/greetings
```

Access HTML UI:
```
$ curl http://localhost:8000/html/index.html
```

UI afficionados may also use a browser of choice. Another nice test client is [Postman](https://www.getpostman.com/)

To test the persitence use the following commands. Mind to replace `<id>`with the id from the initial POST command.
```
$ curl -X POST \
  http://localhost:8000/v1/greetings \
  -H 'content-type: application/json' \
  -d '{
	"Titel" : "Hello",
	"Message" : "Hello, Gopher2!"
  }'

$ curl -X GET \
  http://localhost:8000/v1/greetings/<id>

$ curl -X GET \
  http://localhost:8000/v1/greetings

$ curl -X PUT \
  http://localhost:8000/v1/greetings/<id> \
  -H 'content-type: application/json' \
  -d '{
    "id": "<id>",
	"Titel" : "Hello",
	"Message" : "Hello, again!",
   }'

$ curl -X DELETE \
  http://localhost:8000/v1/greetings/<id>  
```

### Run in Docker ###
There is also a `docker-compose` file so you can run the service in Docker. We are assuming that [Docker CE](https://www.docker.com/) is installed.

To just run a local MongoDB
```
$ docker network create passionator_default
$ docker-compose create passionatordb
$ docker-compose start passionatordb
```


Build container:
```
$ docker-compose build passionator
```

Now create and start the container:
```
$ docker-compose create passionator
$ docker-compose start passionator
```

Logs are avil from:
```
$ docker-compose logs passionator
```

The services can be accessed using the same URLs as when run locally

Just another node: Build, create, and start is posssible in one go (below). For more transparency, I have however chosen to do it step-by-step.
```
$ docker-compose up 
```

To clean up:
```
$ docker-compose stop passionator
$ docker-compose rm passionator
$ docker-compose stop passionatordb
$ docker-compose rm passionatordb
```

### Run in Altemista Cloud ###
Finally, we bring the service into the cloud so that real apps can use it in real live. It is assumed that you have set-up your Altemista Toolbelt and are also familiar with some Altemista Basics. If not I recommend our set of /From Zero to Hero/ tutorials which are avail [here](https://tutorial-tutorial.ballpark.altemista.cloud/).

So we start by logging in to the Altemista Cloud. Replace the parameters iaw with the cluster you are using and your login credentials.
```
$ oc login <cluster>.altemista.cloud:8443 -u "<username>" -p "<password>"
```

Create a new project (with a unique name):
```
$ oc new-project passionator-<your_team_name>
```

Create build credentials so that OpenShift can access Git. Make sure you provide your Git credentials, not the OpenShift credentials. 
```
./createBuildsecret.sh passionator-<your_team_name> <git-user> <git-password>
```

To add persistence on OpenStack based environments run:
```
./createGluster.sh passionator-<your_team_name>
./createAppdb.sh passionator-<your_team_name> passionator gluster test
```

Alternatively, to add persistence on AWS based environments run:
```
./createAppdb.sh passionator-<your_team_name> passionator aws test
```

Finally run to create a test environment for the master branch:
```
./createApp.sh passionator-<your_team_name> passionator https://github.com/Altemista/passionator-golang.git test
```

And here we go:
```
curl -k https://passionator-test-passionator-<your_team_name>.<cluster>.altemista.cloud/html/
```

## Automate build ##
The builds will be automated by adding a web hook to GitLab.

1. Navigate in OpenShift console to Builds / passionator-Test / Configuration and copy the generic web hook URL.
2. Navigate in GitHub to your Project / Settings / Integrations and past the URL in the URL field. Leve secret empty, *uncheck SSL verification* and click "Add webhook".

Click the "Test" button behind the newly created web hook and check that a new build in OpenShift was triggered.

### Create assembly line ###
xxx
yyy
zzz
aaa
bbb
