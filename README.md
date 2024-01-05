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
