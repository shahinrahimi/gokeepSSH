# gokeepSSH

gokeepSSH is a Golang-based utility designed to maintain a persistent SSH tunnel to a remote VPS (Virtual Private Server). The application continuously monitors the SSH tunnel and automatically attempts to re-establish the connection if it terminates.

## Features
- Automatic Reconnection: If the SSH tunnel is terminated for any reason, gokeepSSH will attempt to reconnect after a brief delay.
- Environment Variable Configuration: Configurable through environment variables, making it easy to adapt to different environments and setups.
- Logging: Provides logging for successful connections, errors, and reconnection attempts.


## Requirement
- Go 1.22.4 or later
- A remote VPS with SSH access
- An SSH key for authentication

## Instalation
1. Clone the repository:
```sh
git clone https://github.com/shahinrahimi/gokeepSSH.git
cd gokeepSSH
```
2. Set up the environment variables:
 Create a .env file in the root of the project with the following content:
```sh
VPS_ADDRESS=your_vps_address
LOCAL_PORT=your_local_port
SSH_KEY_PATH=path_to_your_ssh_key
```
3. Build project:
```sh
make build
```

## Usage
Run Application
```sh
make run
```
