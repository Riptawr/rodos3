## Simple library to swtich a rodos-3 relay on or off


## Deps:
- libusb (on linux)
- hidapi
- [GID](https://github.com/karalabe/hid.git)

Tested on linux, no idea whether it runs anywhere else

## Compilation
- compiled with [xgo](https://github.com/karalabe/xgo.git) using --target=linux/arm and --target=linux/amd64

## Usage
- plug in your rodos-3
- start the listener `./rodos3`
- browse to `http://<hostname>:8080/on` or `http://<hostname>:8080/off`

## Troubleshooting
- make sure your user is in the uucp group and the correct usb permissions are added to udev, otherwise run as root
