# Fix requests to internal TLDs in chrome

Often, Chrome thinks some of our hosts are really search strings, sometimes it even issues a search when adding a / after the hosts name.
The only real way to work around this is to use a custom search engine that either redirects to the real search engine or the host url.

## Clone and Build

You need to have the go compiler and the git commandline client installed.

```
$ git clone https://github.com/mbra/fix-chrome-internal-tlds
$ cd fix-chrome-internal-tlds
$ make
```

## Configure and Run

Add a custom search engine with a search string of `http://localhost:8080/?s=%s` and make it the default, now start the fixer:

```
$ ./fix-chrome-internal-tlds -domains=foo,bar,baz
```

Now, for every request that Chrome assumes to be a search term, we will check if thats really the case and either redirect to the search engine or to the detected host.

You may pass a custom search string with `-searchstring` or listen on another ip:port with `-listen`.
