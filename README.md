# Generate Apple Universal links JSON tool

Thanks to the detailed Medium Tutorials

[Universal links in iOS](https://abhimuralidharan.medium.com/universal-links-in-ios-79c4ee038272)

## How to use

```go
link := ulink.NewSimpleUlink(TEAMID,BUNDLEID)
```

Then host it🚀

Your json should look like this

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
