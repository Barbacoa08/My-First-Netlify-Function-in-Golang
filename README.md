# my first netlify function in go

> [Just following these instructions](https://github.com/otaviokr/golang-netlify-example/tree/main) from the Netlify team

Links: [netlify site](https://myfirstgonetlifyfunction.netlify.app/) | [lambda func](https://myfirstgonetlifyfunction.netlify.app/.netlify/functions/hello-lambda)

## dev notes

This does work live, but not locally.

Local error that I have not figured out how to deal with:

```bash
Found incompatible prebuilt function binary in `/functions/hello-lambda.`
The binary needs to be built for Linux/Amd64, but it was built for Darwin/Arm64
```

Run locally: `netlify dev`
> must have Netlify CLI installed: `npm install netlify-cli -g`)
