# Recon
Recon is a application to remotly execute commands on any given target system.

## Security
There is no security implemented at the moment! Use with caution!

## How to use?
It is designed to be used in a docker container.

```bash
docker build --tag recon:$USER .
```

Start the container with
```bash
docker run --user $UID:$GID --name recon -v /path/to/.ssh:/root/.ssh -p 8080:8080 rechon:$USER
```

# Remote controll with recon
General syntax
```
<recon host>:<port>/api/v1/exec/<target host>?cmd="<command>"
```

Shutting down example
```
http GET 192.168.178.5:8081/api/v1/exec/192.168.178.13\?cmd="sudo poweroff"
```
