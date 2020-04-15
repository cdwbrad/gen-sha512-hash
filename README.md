## Compile

To compile ```gen-sha512-hash``` run the following commands:

```
$ git clone https://github.com/rascals77/gen-sha512-hash.git
$ cd gen-sha512-hash
$ go mod download
$ CGO_ENABLED=0 go build -ldflags="-s -w" -o gen-sha512-hash
```

## Execute

To execute the application using command line flags:

```
$ ./gen-sha512-hash -h
Usage of ./gen-sha512-hash:
  -generate
    	generate a random password (optional)
  -length int
    	(default 14) password length of generated password (default 14)
```

To get the SHA512 hash of a password, provide the password when prompted.  The password is masked with asterisks (```*```) as the password is typed. 

```
$ ./gen-sha512-hash
Password: ********
$6$JACJh7kY$7YRb8GTVnvlnGsmu156OqzlmxllxhElo.aIU1sOqf7tSQ9ETcWJN/cZfOq2O5sdMwafApvttGMZdvHxcuAnCK0
```

To generate a random complex password and it's corresponding SHA512 hash, use the ```-generate``` flag:

```
$ ./gen-sha512-hash -generate
G_w?cWFu3Aj0nw
$6$yQ7uQTVY$7giRTvvqh90GZiX0VwhWv93qo7fUT9hlifx8IAcBk9uE62VUWx5fSVJjnN6QCmYHoR5yV78yuUMZWn2MUFR7D.
```


To generate a random complex password with a given length, use the ```-length``` flag along with the ```-generate``` flag:

```
$ ./gen-sha512-hash -generate -length 20
ob.8drOyWwIF!V&)b`%V
$6$BtATwcyA$NHl6Otn1fYh7SmLdJa6ewlDbDHH58uZu3dPzKavbecNs5q1qXykFW9Kx8KpwQWMeX03oa6qzhES8srT7ijHV51
```
