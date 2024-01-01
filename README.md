# file-serve

A simple HTTP server that serves a given file from a given path.

```bash
./file-serve -path /releases/v0.1.0/my-tarball.tar.gz -file ./my-tarball.tar.gz
```

You can then run `curl -O http://localhost:8888/releases/v0.1.0/my-tarball.tar.gz` and download the file.
