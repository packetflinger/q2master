# q2master
This project is a custom implementation of a Quake 2 master server written in Go. 

## Command line arguments

All configuration options for this server are handled through the command line arguments:

`-port <###>`  
Specify the listen port for Quake 2 gameservers. The default is `27900` and probably should not be changed unless there are very specific circumstances and you know what you’re doing.

`-addr <x.x.x.x>`  
This is the IP address to listen on. By default the server will listen on all IP addresses of the host machine, use this to specify a single address. Default is `[::]` for all IPv4 and IPv6 addresses. You can use something like `0.0.0.0` for all IPv4-only addresses.

`-fg`  
Run the server in the foreground. This will output all log information to the console where the server was run and will not write to a log file.

`-logfile <filename>`  
Specify the path and name of the file to use for logging. The default is `master.log` in the same directory where the server was started from.

`-api`  
Enable the built-in HTTP server for API requests. Default is disabled

`-apiport`  
The TCP port to listen on for API requests. Default is `3333`

`-apiaddr`  
The IP address to listen on for API requests. Default is `[::]` for all addresses.

## API endpoints

`/GetServers`   
Fetch all the Quake 2 servers reporting to this master server

`/HealthCheck`  
Get some basic information about the server

`/ServerInfo`  
Get the data for a particular server. Requires a `srv` query string parameter which is the IP:port of the game server.
