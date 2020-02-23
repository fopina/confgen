# confgen
Generate configuration files from Go templates

As the (internet) world is all going serverless and [12-factor](https://12factor.net/) for scalability and resilience and such, `envsubst` became quite popular to "12-factor" some apps that only support configuration through files.

Isse with `envsubst` is that it only does... env substitutions, nothing else. No logic control to allow for more complex templating (to eventually simplify/minify the required env vars).

An easy (and probably popular) option is to use [Jinja2](https://palletsprojects.com/p/jinja/) templates.  
(python) Developers already know it from Django/Flask, devops know it from ansible.

But for a python-less image (such as nginx), adding ~50M of python runtime just to have templating sounds a bit too much.

`confgen` tries to help with that.
Instead of Jinja2 templating, it's [Go templating](https://golang.org/pkg/text/template/).
Instead of 50M footprint, it's 2M.

Currently it simply injects the `env` function into the templates mapped to Go `os.Getenv`.

## Usage

* Add confgen in your Dockerfile

```
...
RUN wget https://github.com/fopina/confgen/releases/latest/download/confgen_linux_amd64 -O /usr/local/bin/confgen
RUN chmod a+x /usr/local/bin/confgen
...
```

* Call it from your entrypoint

```
...
confgen -o /path/to/used/conf /path/to/template
...
```

* With your [template](https://golang.org/pkg/text/template/) using the `env` function

```
config_key={{ "MY_CONFIG" | env }};
```

* Set those env vars when creating the containers

```
docker run -e MY_CONFIG=wtv ...
```

Check [examples](examples) for actual usage with nginx paired with `envsubst`. And check [Go template](https://golang.org/pkg/text/template/) documentation for full list of logical controls.
