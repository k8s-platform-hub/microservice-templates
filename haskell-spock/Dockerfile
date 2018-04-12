FROM haskell:8.0.2

RUN mkdir /usr/src/app

COPY haskell-spock.cabal /usr/src/app/haskell-spock.cabal
COPY stack.yaml /usr/src/app/stack.yaml

WORKDIR /usr/src/app

RUN stack --stack-yaml /usr/src/app/stack.yaml setup
RUN stack build --only-dependencies

COPY . /usr/src/app

RUN pwd
RUN stack build --fast
RUN stack install --local-bin-path /usr/bin
RUN rm -rf /usr/src/app /opt/ghc /root/.stack
CMD haskell-spock
