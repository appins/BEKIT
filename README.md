[![AppIns' BEKIT](https://imgur.com/zfCL7nD.png)](https://github.com/appins)

## A tool for creating web servers with ease.
BEKIT stands for Backend Kit.

### Getting Started
Let's get started! Just open a terminal and type:
```sh
  git clone ...
  cd bekit/src
  go build -o ../bin/bekit && ../bin/bekit
```
 The BEKIT Command Line Tool should now be running!

### How To Use
To start using Bekit, you must first know what port you want
to use for your web server. Some ports will require you start
Bekit with sudo, or root, privileges.

If you want to start your server on port 80, you would type
`onport 80` and then hit enter. To end this particular block,
you would type `end` and hit enter. This will automatically start
the server (This implementation **will** change in the near future).

Inside of each block, there are multiple statements to customize how
your server will run. Every block needs at least one `main` folder.
All of the following commands need to be executed in a block.

The `main` folder declares which folder the web server runs from.
To set up the `main` folder, just type `main /path/to/folder`. This
folder must contain a file called index.html, even if you don't plan
on using it.

You can reroute specific files easily. If a user wants to access a file called
`contact.html`, you can reroute them to `about-us.html` by administrating the
command `filerr contact.html->about-us.html`. If the rerouted file starts with
a `/`, it will interpret it as a file somewhere else on your hard drive.
