Tuner
=====

A small configuration management system for updating performance config in dynamic systems.

## What this solves

When building Docker containers we locking in a set performance configuration. This means that
you have to either deploy the container at a set size (CPU/Memory) OR write some "interesting"
scripts to configure your environment before it boots.

This project goes with the latter, but handles it in a clean way.

Configuraion is handled by environment 3 variables:

* `TUNER_MAX` - The maximum amount of memory available
* `TUNER_PROC` - The max memory per process
* `TUNER_MULTIPLIER` - A method of boosting how many procs can be run (over provisioning)

Configuration is them returned via inbuilt templates to Stdout. You then have the option to pipe
this configuration to a known local on your host (prior to the process booting) eg.

```bash
$ tuner --conf=apache > /etc/apache2/mods-enabled/tuner.conf
```

## Install

```bash
$ curl -L https://github.com/previousnext/tuner/releases/download/1.0.0/tuner-linux-amd64 -o /usr/local/bin/tuner
$ chmod +rx /usr/local/bin/tuner
```

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

$ export TUNER_PROC=256
$ tuner --conf=php

memory_limit = 256M
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
