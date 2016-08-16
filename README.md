Tuner
=====

Configures system based on 3 variables:

* Max - The maximum amount of memory available
* Proc - The max memory per process
* Multiplier - A method of boosting how many procs can be run

## Usage

**Apache**

```bash
$ tuner --conf=apache
<IfModule mpm_prefork_module>
	StartServers		2
	MinSpareServers		2
	MaxSpareServers		2
	MaxRequestWorkers	8
	MaxConnectionsPerChild  8
</IfModule>
```

**PHP**

```bash
$ tuner --conf=php
memory_limit = 128M
```

## Advanced usage

```bash
$ export TUNER_CONF=apache
$ export TUNER_MAX=1024
$ export TUNER_PROC=64
$ export TUNER_MULTIPLIER=3
$ tuner --conf=apache
<IfModule mpm_prefork_module>
	StartServers		2
	MinSpareServers		2
	MaxSpareServers		2
	MaxRequestWorkers	48
	MaxConnectionsPerChild  1024
</IfModule>
```

## Contributing new conf plugins

* Create a new folder with the machine name of your plugin in:

```
src/github.com/previousnext/tuner/conf/NEW
```

* Create a `conf.go` file in this directory which adheres to the inferface in:

```
src/github.com/previousnext/tuner/conf/conf.go
```

Note: See the PHP conf for a simple example.

* Create a test to ensure the correct values are built in the template.

```
src/github.com/previousnext/tuner/conf/NEW/conf_test.go
```

Note: See the PHP conf for a simple example.

* Ensure that the new plugin is added to the `main.go` import statement like the others.
