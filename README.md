# Generate Apple Universal links JSON tool

Thanks to the detailed Medium Tutorials

[Universal links in iOS](https://abhimuralidharan.medium.com/universal-links-in-ios-79c4ee038272)

## How to use

```go
link := ulink.NewSimpleUlink(TEAMID,BUNDLEID)
```

Then host it🚀
It should be reached with:

```text
NAME.YOUR.DOMAIN/apple-app-site-association
or
NAME.YOUR.DOMAIN/.well-known/apple-app-site-association
```

Your json response should look like this

```json
{
  "applinks": {
    "apps": [],
    "details": [
      {
        "appID": "JHGFJHHYX.com.facebook.ios",
        "paths": [
          "*"
        ]
      }
    ]
  }
}
```

## Test your domain

[VALIDATOR]( https://branch.io/resources/aasa-validator/#resultsbox)