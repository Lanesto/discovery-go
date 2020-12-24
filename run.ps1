$image = "discovery-go:local"

docker image build -t $image .
echo ""

docker container run -it --rm `
    -v ${PWD}:/workspace `
    -p 8080:8080 `
    -w /workspace `
    $image /bin/bash

pause