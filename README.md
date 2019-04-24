# Bitmex-status-cli

![release 1.0](https://img.shields.io/badge/release-1.0-green.svg)

Get your positions and orders status with a simple command. 


#### Minimum Specifications

- Go 1.10 : https://golang.org/doc/install.
- Tested on Windows 10, Linux 18.10

#### Setup
``cd ~/go/src/``

``git clone https://github.com/2lazy2debug/bitmex-status-cli``

``cd bitmex-status-cli``

##### Dependencies
``go get github.com/fatih/color``

##### API Key
You should put your API key in a bitmex.apikey file on the same level as main.go. If you decide to fork the project, there's already a reference in .gitignore.

Since I'm doing this for personal purpose, I'm not encrypting shit (api calls are in https of course) so... careful

Syntax : 

``id:myverypersonalbitmexid``

``secret:myveryveryverypersonalbitmexsecret``

``Add a line break``

##### Run 
``go run main.go``

#### Screenshot
![Screen1](https://github.com/2lazy2debug/bitmex-status-cli/blob/master/bitmex-screenshot1.png?raw=true)

#### Referral 
If you haven't joined BitMEX yet (why are you here then?) and you want 10% discount on fees : https://www.bitmex.com/register/1TTdTj


