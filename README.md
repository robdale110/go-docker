## Go Docker Development Environment

Can build using either

```docker build --target bin --output bin/ --platform local .```

or 

```make bin/example```

```make```

build for specific platform

```docker build --target bin --platform windows/amd64 .```

use make for build, test and lint:

```make unit-test```
```make lint```