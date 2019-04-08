# Bitmex-status-cli

##### Work in progress - more to come SoonTM

Actual situation : You just execute go run main.go and get the raw position response for the server. 

Next steps:

- Formatting returned data nicely
- Return stop and profit orders as well
- Integration for Travis
- Testing


##### Minimum Specifications

- Go 1.10 : https://golang.org/doc/install.
- Tested on Windows 10, Linux 18.10

##### Setup
``cd ~/go/src/``

``git clone https://github.com/2lazy2debug/bitmex-status-cli``

``cd bitmex-status-cli``

``go run main.go``

##### API Key
You should put your API key in a bitmex.apikey file on the same level as main.go. If you decide to fork the project, there's already a reference in .gitignore.

Since it's kind of my first project in Golang, I'm not encrypting shit (api calls are in https of course) so... careful

Syntax : 
``id:myverypersonalbitmexid``
``secret:myveryveryverypersonalbitmexsecret``
``Add a line break``

##### Referral 
If you haven't joined BitMEX yet (why are you here then?) and you want 10% discount on fees : https://www.bitmex.com/register/1TTdTj


