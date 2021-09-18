![echo.go - the simple out-of-memory exception generator](res/logo.png)

echo.go 💨 is an out-of-memory exception generator written in go 💨 as a companion to [hello-world.rs](https://github.com/mTvare6/hello-world.rs.git). since it's in go 💨 the code is simpler, shorter, and easier to read and portable and good package management compared to c(rap) or rus(hi)t! binaries are only 1.9 million bytes each 💨

## Installation

echo.go 💨 can be installed by downloading the most recent source code release and choosing the appropriate binary for your operating system and architecture and copying it to the appropriate application directory 💨

## Building from source

Obtain the source with a git clone or download a release tarball. To use the deployment script you will need to have go 💨 and bash with commands echo, cut, mkdir. If you just want to build for your OS you just need go 💨

### Cross-system deployment

A script is provided, build.sh to invoke go 💨 build for all the available target architectures. make sure you are in the project root directory:

    ./build.sh

Binaries (only 1.9 MB each, a total of just ~82 MB 💨) will be located in the dist/ folder.

### Local build

    go 💨 build -o echo src/main.go 💨

## Usage

    (arch)-echo [args]

The program will then ignore your arguments and output a long memory exception trace 💨