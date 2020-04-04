# sshp


## What?
A fake shell that sends user input in a Discord channel using a webhook. I made this primarily to setup a honeypot on my server by changing the port of the main SSH service and starting a containerised SSH service on port 21 with a passwordless root user, where the shell of `root` points to this fake shell. Setting up a honeypot can help you find possible vulnerabilities in your system.

## How (to use this)?
- Install [Go](https://golang.org/dl/) and [Make](https://www.gnu.org/software/make/)
- Clone this repository
```sh
$ git clone https://github.com/y21/sshp
```
- Use the given Makefile to build the binary
```sh
$ make shell
```
- For convenience, you can create a symbolic link in `/bin/`
```sh
ln shell /bin/shell
```
- Edit the `config.json` file

You can now run the fake shell. 
> Only follow the next points if you know what you are doing. If you don't, you might end up losing access to your machine.
If you decide to change the login shell of `root`, make sure to create another superuser so you can always change it back.
Feel free to create a containerised SSH service for this
- Change the login shell of the current user
```sh
chsh -s /path/to/fake/shell
```


## TODO
- Send IP Address of attacker
- Add more fake commands
- Restructure source code