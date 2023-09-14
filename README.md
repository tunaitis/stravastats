# stravastats

stravastats is a command line utility to display your personal Strava statistics in the terminal.

## Installation

There are a few ways to install stravastats for each platform. 

```bash
# Homebrew (macOS & Linux)
$ brew tap tunaitis/tools
$ brew install stravastats

# Windows
$ scoop bucket add org https://github.com/tunaitis/stravastats.git
$ scoop install tunaitis/stravastats

# Using Go directly
$ go install github.com/tunaitis/stravastats@latest
```

You can also download and manually install the binary from the [release page](https://github.com/tunaitis/stravastats/releases).

## Setup

stravastats need to authorize with Strava API before receiving any athlete data. The API uses the OAuth authorization protocol, and specifically the code authorization flow, to grant access to third parties. The code authorization flow is carried out in the following steps:

* stravastats opens a browser window to send the user to the Strava OAuth server
* The user sees the authorization prompt and approves the access request 
* The user is redirected back to the stravastats' internal web server with the authorization code in the query string
* stravastats exchanges the authorization code for an access and refresh tokens

The tokens are then placed in a secure vault and used to access the athlete's personal data later. 

### Create a Strava App

Sign into your Strava account and go to [https://www.strava.com/settings/api](https://www.strava.com/settings/api). Once there, fill the form with the appropriate values and set the authorization callback domain to localhost. Copy the application id and secret, which you will need in the next step.

### Set application configuration

Run the following command to initialize the configuration. 

```bash
$ stravastats init
```

When prompted, enter the application id and secret.

### Authenticate
