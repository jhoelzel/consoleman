# Consoleman

```Console
------------------------------------------------------------------------	
_________                            .__                                
\_   ___ \  ____   ____   __________ |  |   ____   _____ _____    ____  
/    \  \/ /  _ \ /    \ /  ___/  _ \|  | _/ __ \ /     \\__  \  /    \ 
\     \___(  <_> )   |  \\___ (  <_> )  |_\  ___/|  Y Y  \/ __ \|   |  \
 \______  /\____/|___|  /____  >____/|____/\___  >__|_|  (____  /___|  /
        \/            \/     \/                \/      \/     \/     \/ 
			Like Postman, but in the console!	
------------------------------------------------------------------------		
```

Consoleman is a command-line utility that acts like Postman but runs in the console. You can use it to send HTTP requests to APIs and inspect the responses.

Disclaimer: I am aware that this does not much more than curl, but it is an exercise using ASCII interfaces with golang and to prove the point that it is possible without dependencies.

## Installation

### Using `go install`

To install Consoleman using `go install`, first clone the repository or download the source code. Make sure you have a Go module set up for your project (refer to the `go.mod` file).

Then, run the following command in your project's root directory:

```Console
go install 
```

or simply run

```Console
go install github.com/jhoelzel/consoleman
```

This command will build and install the Consoleman binary to your `$GOPATH/bin` directory. Make sure that `$GOPATH/bin` is in your system's `$PATH` to access the Consoleman binary from anywhere in your system.

### Building the binary manually

Alternatively, you can build the Consoleman binary manually. To do this, run the following command in your project's root directory:

```Console
go build -o consoleman main.go
```

Now you can move the `consoleman` binary to a directory in your `$PATH` to make it accessible from anywhere in your system.

## Usage

You can use Consoleman interactively or by providing command-line flags.

### Interactive Mode

Run Consoleman without any flags to use the interactive mode:

```Console
./consoleman
```

<this is where i would put my gif: If I would have made one yet ;)>

In interactive mode, Consoleman will prompt you for the following information:

1. Protocol (http or https)
2. URL
3. Request Type (GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS, CONNECT)
4. Basic Auth credentials (username:password) - optional
5. Headers (Key:Value, separated by ';') - optional
6. Body - optional

### Command-Line Flags

You can also provide the request details as command-line flags:

```Console
./consoleman -protocol=https -url=example.com -requestType=GET -auth=username:password -headers="Content-Type:application/json" -body='{"key": "value"}' 
```

The available flags are:

```Console
- `-protocol`: Protocol (http or https)
- `-url`: URL
- `-requestType`: Request Type (GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS, CONNECT)
- `-auth`: Basic Auth credentials (username:password) - optional
- `-headers`: Headers (Key:Value, separated by ';') - optional
- `-body`: Body - optional
- `-noUI`: Skip empty optional flags when any flags are provided and hide interface
```

### Skipping Empty Flags

If you want to skip interactive prompts for any empty flags, use the `-noUI` flag:

```Console
./consoleman -protocol=https -url=example.com -noUI
```

In this example, the `requestType`, `auth`, `headers`, and `body` flags are not provided, and since the `-noUI` flag is set, their respective interactive input prompts will be skipped.

## Response Output

After sending the request, Consoleman will display the response body. If there was an error with the request, an error message will be displayed instead.

## Examples

Here are some example use cases of Consoleman:

1. Send a simple GET request to a public API:

```Console
./consoleman -protocol=https -url=api.example.com/users -requestType=GET
```

2. Send a POST request with a JSON body and custom headers:

```Console
./consoleman -protocol=https -url=api.example.com/users -requestType=POST -headers="Content-Type:application/json;Authorization:Bearer my_token" -body='{"name": "John Doe", "email": "john@example.com"}'
```

3. Send a PUT request with Basic Authentication:

```Console
./consoleman -protocol=https -url=api.example.com/users/1 -requestType=PUT -auth=my_username:my_password -body='{"email": "new-email@example.com"}'
```

4. Send a DELETE request:

```Console
./consoleman -protocol=https -url=api.example.com/users/1 -requestType=DELETE
```

5. Using the `-noUI` flag to skip input prompts for empty flags:

```Console
./consoleman -protocol=https -url=api.example.com/users -requestType=GET -noUI
```

In this example, the `auth`, `headers`, and `body` flags are not provided, and since the `-noUI` flag is set, their respective interactive input prompts will be skipped.

## Contributing

If you would like to contribute to the development of Consoleman, feel free to fork the repository and submit a pull request with your changes. Please follow the existing code style and include comments when necessary.

## License

Consoleman is released under the MIT License. See the `LICENSE` file for details.
