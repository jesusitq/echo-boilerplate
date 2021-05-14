# Boilerplate
Api base

## Getting started

In order to run this API locally, please make sure you have the next dependencies:

    go version 1.11.4 or above

And ensure that the `go` command is available in the path by running the command:

    go version

Then, clone the repository in your `GOPATH`.

### Initial Configuration

To run the API you will need to download the project libraries:

  * `go get ./...`

### Building

In order to build the API, run the next command:

    go build

This will generate an executable file according your operating system, for example for GNU/Linux:

    ./boootstrap

Or for example for Windows:

    .\boootstrap.exe

### Running the API

To run the API, you need to set the next environment variables:

#### `ENV`

Set the environment to run the API. Can have the next values:

 - `dev` for development environment (by default).
 - `prod` for production environment.

#### `LOG_LEVEL`

Set the log level to print. Will print the value and above. Can have the next values:

 - `0` for debug.
 - `1` for info.
 - `2` for warning (by default).
 - `3` for error.
 - `4` for critical.

One you set the environment variables, you can run the next command in order to start the API:

    go run main.go

Or if you compiled the project in the previous step, use the command to run the executable according your operating system.

If you see the next output in the console, the API has been started successfully:

    INFO [main] [main] Started <bootstrap> in <dev> environment on port <4001>


## API endpoints

This section describes the calls that accepts the API


### `POST /todo/graphql/`

Endpoint example
