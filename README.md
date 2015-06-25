World-Generation
======

Setup
-----

We are using Vagrant for development. When you run the following command a new
box will be provisioned. Go 1.4 will be installed and a go workbench will be
setup.

```
  vagrant up
```

Running
-------

```bash
  # Make sure you are in your vagrant box
  vagrant ssh

  # Build and run the app
  go build && ./world-generation

  # Visit maps to randomly generate a map
  curl localhost:8080/maps?x=1&y=3
```

_note: vagrant has been setup to forward the default port (8080) to your host
machine allowing you to access everything from your web browser_
