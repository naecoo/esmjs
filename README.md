**One Command transform Javascript to ESM module**


# Install

## Git

```bash
git clone https://github.com/naecoo/esmjs.git

./esmjs/bin/esmjs --help

```

## HomeBrew

```bash
brew tap

brew install

esmjs --help
```


# Usage
```
esmjs input.js output.js

# typescript
esmjs input.ts output.js

# minify
esmjs input.js output.js -m

# build for browser platform
esmhs input.js output.js -m -b
```


## Roadmap

- ~~file struct define~~
- ~~import esbuild~~
- ~~define CLI opiton and flag (features?)~~
- ~~coding~~
  - ~~basic functions~~
  - ~~log, signal and erros~~
- ~~test and benchmark~~
- ~~example~~
- ~~docs~~
- publish
