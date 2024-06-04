ARG tag=latest
ARG baseimage=etherspot/skandha

FROM $baseimage:$tag
ARG TARGETPLATFORM

RUN echo "I'm building for $TARGETPLATFORM"


ADD skandha_node.sh /skandha_node.sh
RUN chmod +x /skandha_node.sh

RUN node --experimental-specifier-resolution=node ./packages/cli/bin/skandha.js --version | grep "Version:.*" | grep -o "v.*" > /version.txt

ENTRYPOINT ["/skandha_node.sh"]
