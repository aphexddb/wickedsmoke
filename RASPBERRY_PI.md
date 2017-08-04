# Raspberry Pi setup

This setup guide uses [this guide](https://learn.adafruit.com/raspberry-pi-zero-creation/overview). Download the [Raspbian lite](https://downloads.raspberrypi.org/raspbian_lite_latest) image. If you are using [console cable](https://learn.adafruit.com/raspberry-pi-zero-creation/give-it-life) pay attention to what pins are power!

Use a nice fast memory card, not a cheap one.

## OS and basic setup

1. Flash image to disk
    ```bash
    sudo dd bs=1m if=/Users/USERNAME/Downloads/2016-11-25-raspbian-jessie-lite.img of=/dev/rdiskn conv=sync
    sudo dd bs=1m if=/Users/USERNAME/Downloads/2016-11-25-raspbian-jessie-lite.img of=/dev/rdisk3 conv=sync
    ```  
2. Add empty file `ssh` on root of `/boot` raspbian disk image to enable ssh
3. Add to bottom of `/boot/config.txt` to enable GPIO serial via FDTI
    ```bash
    enable_uart=1

    # Have not tried this yet
    #dtparam=i2c1=on
    #dtparam=i2c_arm=on
    #dtparam=spi=on
    ```
4. Create `/boot/wpa_supplicant.conf` (must have unix style LF line endings)
    ```
    network={
      ssid="your_ssid_here"
      psk="password"
      scan_ssid=1
    }
    ```
5. Plug Pi into USB (now power!) and login after waiting ~1.5m, password is `raspberry`
    ```bash
    ssh pi@raspberrypi.local
    ```

## I2C Bus

1. SSH in, enable I2X kernel module, enable Serial bus and enable SPI.

    ```bash
    sudo raspi-config
    ```

2. Reboot / power cycle

3. Install I2C tools

    ```bash
    sudo apt-get install -y i2c-tools
    ```

4. The `i2cdetect` program will probe all the addresses on a bus, and report whether any devices are present.

    ```bash
     i2cdetect -y 1
    ```    

## Testing

```bash
sudo apt install -y python-gpiozero
```

```python
from gpiozero import MCP3008
from time import sleep
 
while True:
    for x in range(0, 8):
        with MCP3008(channel=x) as reading:
            print(x,": ", reading.value)
sleep(0.1)
```