# twitter credential maker
make twiter credential with anaconda

##example
```go
// set consumer key and secret...
c, err := cred.New()
if err != nil {
    // error handle
}
api := anaconda.NewTwitterApi(
    c.AccessToken,
    c.AccessSecret,
)
```


##thanks
[anaconda](https://github.com/ChimeraCoder/anaconda) by ChimeraCoder  
[go-homedir](https://github.com/mitchellh/go-homedir) by mitchellh
