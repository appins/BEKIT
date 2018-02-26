[![#AppIns' BEKIT](https://imgur.com/zfCL7nD.png)](https://github.com/appins)

## A tool for creating web servers with ease.
BEKIT stands for Backend Kit.

### Getting Started
Let's get started! Just open a terminal and type:
```sh
  git clone ...
  cd bekit/src
  go build -o ../bin/bekit && ../bin/bekit
```

 The BEKIT CLI should now be running.

### Alpha Notice:
Each update will add or change existing features. I can not
ensure that an update will not break your code (for now).

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

The bare minimum for a BEKIT server is a port and a main folder,
a server like that would look something like
```
onport 80
main /path/to/folder
end
```

You can reroute specific files easily. If a user wants to access a file called
`contact.html`, you can reroute them to `about-us.html` by administrating the
command `filerr contact.html->about-us.html`. You can also disable access
to a file by rerouting it to `null`, such as
`filerr contact.html->null`.

The `logip` command will log IP's of users who request files that
do not exist. This is useful to see who may be trying to maliciously
request files.

The `force` command will ignore most errors and run. This should be put
at the very top of your code. Even with the force command, you still
need to define a main folder. The `force-lite` command will only ignore
small errors such as commands not being recognized. These commands
have opportunities to cause MAJORS issues if the command has incorrect
syntax. Don't use either except for quick testing.

Commands will continue to be added until the release stage. The current
stage is: `alpha`.

### Sample Usage:
If we wanted to start a server on port 80 with root path of /server/www/
and we want to log IP's, you would start the BEKIT CLI and enter in:
```
> onport 80
    { main /server/www/
    { logip
    { end
```
(The `>` and `{` are already on the command line for you, do not enter those
  in yourself)
