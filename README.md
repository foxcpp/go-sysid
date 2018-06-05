go-sysid
==========

Why?
------

Some vulnerabilities allow reading victim's files but not more. Using similar
holes we can steal cookies, keys, session information and other confidential
information. What we can do to protect our software from similar
vulnerabilities in 3rd party software found on user's machines? The answer is
simple: Encrypt everything. But people often dislike requirement to use
passwords everywhere, so we need a way to generate unique key per machine and
use it... Right! Machine, use machine information. A lot of machine
information can't be gathered without RCE vulnerability and a lot of
information is reasonably unique. 

This is not strong protection, my personal recommendation is to use
passwords if possible.


**Note:** Only Linux is currently supported, Windows, Mac OS may get support in future.

Usage
-------

Install library using `go get`:
```
go get github.com/foxcpp/go-sysid
```

You can also grab console utility here:
```
go get github.com/foxcpp/go-sysid/sysid
```

In code you just call `SysID()` and it works. If you need some tuning (for
example, use different hash function because you need different key size)...
See `SysIDCustom()`.


There are two build tags for console utility: `blake2b` and `sha3` to enable
additional hash functions support.


Security issues
-----------------

`fox.cpp at disroot dot org`. Use PGP encryption if possible, here is my key fingerprint:
```
3197 BBD9 5137 E682 A597
17B4 34BB 2007 0813 96F4
```

License
---------

MIT.
