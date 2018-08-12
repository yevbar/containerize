# containerize

Made at the Stupid Hackathon

Ever needed to nest docker containers inside docker containers endlessly? Well now you can!

## why

Several companies run applications inside docker containers which is great until a malicious actor is able to break out of the container and infiltrate the host server.

The solution, put our container with the application inside another container inside another container and so on till you see fit. If applied properly, the application will be so deep the attacker will have to ask philosophical questions.

## how to use

Build the go file then run (with root permissions for keylogger)

```bash
$ go build containerize.go
$ sudo ./containerize <Number of iterations> <command> [rest of command]
```

The first argument provided is the number of containers deep you want your command/script to run

_**example**_

```bash
$ go build containerize.go
$ sudo ./containerize 4 echo "hi"
```

This should take you into a container inside a container inside a container inside a container and then run `echo "hi"` once there
