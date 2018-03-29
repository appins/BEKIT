[![#AppIns' BEKIT](https://imgur.com/zfCL7nD.png)](https://github.com/appins)

# A tool for creating web servers with ease.
### BEKIT stands for Backend Kit.

### Getting Started
Let's get started! Just open a terminal and type:
```sh
  git clone https://github.com/AppIns/BEKIT.git
  cd bekit/src
  go build -o ../bin/bekit && ../bin/bekit
  # Or, to run as root, which is required for all ports under 1024
  go build -o ../bin/bekit && sudo ../bin/bekit
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

# Block opening commands:
These commands must be typed before any others, they open a block of BEKIT code.

## `> onport` port
Determines what port your server is running on. Stats it right afterwards. This is
the cli method of using BEKIT

## `> save` filename
Opens a block of code that is saved. It can later be ran by recalling it using
the `load` command. Code is saved in a format of how it was typed.

If the file does not contain a '.' then it will automatically be set up with the
extension .beks (backend kit save). This also applies to loading files.

## `> load` filename
Opens and loads a file that was saved.

## `main`
The `main` folder declares which folder the web server runs from.
To set up the `main` folder, just type `main /path/to/folder`. This
folder must contain a file called index.html, even if you don't plan
on using it.

The bare minimum for a BEKIT server is a port and a main folder,
a server like that would look something like:
```
onport 80
main /path/to/folder
end
```

## `filerr`
You can reroute specific files easily. If a user wants to access a file called
`contact.html`, you can reroute them to `about-us.html` by administrating the
command `filerr contact.html->about-us.html`. You can also disable access
to a file by rerouting it to `null`, such as
`filerr contact.html->null`.


#### Example Usage
The following will take all requests to main.html and reroute them to index.html
```
filerr main.html->index.html
```

## `logip`
The `logip` command will log IP's of users who request files that
do not exist. This is useful to see who may be trying to maliciously
request files or overload the server.

## `force` and `force-lite`
The `force` command will ignore most errors and run. This should be put
at the very top of your code. Even with the force command, you still
need to define a main folder. The `force-lite` command will only ignore
small errors such as commands not being recognized. These commands
have opportunities to cause MAJORS issues if the command has incorrect
syntax. Don't use either except for testing.

## `f`
The `f` command is experimental and incomplete. It takes data from the
server and then copies it to somewhere else. You can take data from
requests and 404 errors **for now** and you can only output to console **for now**. More stuff will be added soon!

The syntax for the `f` command is `f ROOT:METHOD->ROOT:METHOD`.
The left side of the '->' is input, while the right side is
output. An example of a root input would be "request" while an example
root output would be "write". Methods are specific to what you are
using, and usually act as parameters. Bellow is a list of available roots and methods.

The following keywords are available for the left side:
  * `request` (methods are: `:ip`, `:file`, and `:is404`)
  * `404` (methods are: `:ip` and `:file`, and `:is404`)
  * `form` (the method is the value of the form, such as `:name`)
  * `text` (the method is the text that you want to pass to the output, such as `:hello world`, `:newline` will be replaced with a newline)

The following keywords are available as outputs
  * `console` (takes `:nonewline` as a possible method)
  * `write` (which takes a filename as a method, newlines will have to be added using `text`)

#### Example Usage
The following code will write users requests, such as "/index.html x.x.x.x" to the console
```
f request:file->console:nonewline
f text: ->console:nonewline
f request:ip->console
```

## `#` and `//` comments
If a line starts with `# ` or `// `, the line will be ignored.
There must be a space after the `#` or `//`.

Commands will continue to be added until the release stage. The current
stage is: `alpha`.

### Future Ideas
BEKIT is being developed actively. For now, the basics are still being
implemented. We should see these goals met soon.

  * A cloning feature where one server can be created as an exact clone
of another. The server being cloned will have to enable it and a file on
the server will be used as the clone.
  * Inputting files instead of just hand-written code. Could save some
major time
  * Output and input of .json files (should be implemented before cloning)
  * Outputting go code for your server.


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
