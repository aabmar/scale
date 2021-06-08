
Reading DYMO M5 data with Gopher Interface Devices
==================================================

This is a smal Go program reading data from a DYMO M5 scale and prints out the number of grams on the console.

The program uses the [Gopher Interface Devices (USB HID](https://github.com/karalabe/hid). The github.com/karalabe/hid git module does not have any examples or tutorials included. This program can be used as a simple example on how to use it to read data from a HID device.


Installation
============

Make sure you have golang 1.15+ installed on your computer.

```
git clone https://github.com/aabmar/scale
make init
make run
```

It is only tested on MacOS.	

