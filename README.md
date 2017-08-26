# GRPC Phonebook

# Setup

### Go

This requires [govendor](https://github.com/kardianos/govendor) to work properly. You can also install the server by doing `go install github.com/iheanyi/grpc-phonebook/cmd/pbsrv` and then running `pbsrv`.

### Node.js

Install dependencies using `npm` or `yarn`. You can run the command line client by doing `./cmd/pbctl/pbctl [options] <command>`.

```
  Usage: pbctl [options] [command]


  Options:

    -V, --version  output the version number
    -h, --help     output usage information


  Commands:

    create [options] <name>
    list                   
    show <name>            
    delete <name>          

```

Valid options are:

```
-e, --email <email>  email for the contact
-w, --work   <number> work number for the contact
-h, --home   <number> home number for the contact
-m, --mobile <number> mobile number for the contact
```
