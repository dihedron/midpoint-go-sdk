# MidPoint Golang Software Development Kit

## How to install MidPoint locally for testing

Download the setup script and launch it:

```bash
$> mkdir midpoint && cd midpoint
$> curl -O https://raw.githubusercontent.com/Evolveum/midpoint-docker/master/midpoint-quickstart.sh
$> chmod 744 midpoint-quickstart.sh
$> ./midpoint-quickstart.sh start
```

When requested, insert a test password (8 characters minimum, including at least a number, a lowercase and an uppercase character), e.g. `S3cret!!`.

To forward the port when remotely connected via VSCode SSH Remote mode, navigate to the the Ports tab and add the 8080 port forwarding.

To stop the running container and destroy associated resources:

```bash
$> ./midpoint-quickstart.sh delete
```

## How to call REST APIs from the command line

Tu make a call with JSON payload and result (chaange the password as needed):

```bash
$> curl --user 'administrator:S3cret!!' -X GET http://localhost:8080/midpoint/ws/rest/self -H "Accept: application/json" -H "Content-Type: application/json"
```

