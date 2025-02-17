# Crestron DM Switcher
A switcher maps one of multiple input signals (audio and/or video) to a specified output.  This microservice communicates with and controls the Crestron switchers via Telnet.
It works with the following Crestron switchers, even though only the DM-MD8x8 was tested:

DM-MD8x8

DM-MD16x16

DM-MD32x32

[Crestron DigitalMedia Matrix Switchers](https://www.crestron.com/Products/Featured-Solutions/DigitalMedia-Matrix-Switchers)

![](https://github.com/Dartmouth-OpenAV/microservice-crestron-dm-switcher/blob/main/front.png)
![](https://github.com/Dartmouth-OpenAV/microservice-crestron-dm-switcher/blob/main/rear.png)

The following operations are supported:  
set **'avroute'**, set **'videoroute'**, get **'avroute'**, get **'videoroute'** using the same endpoint syntax as Kramer switchers.  
**'avroute'** and **'videoroute'** are essentially the same - audio follows video.

The below bash script was used for stress testing the DM-MD8x8.  Over 100 loops were run without errors.

```bash
#!/bin/bash

SWITCHER=<IP ADDRESS>
sleep 2
let count=0
while true
do
    date
    let count=count+1
    echo "**** loop: " $count "****"

    curl http://localhost:80/${SWITCHER}/videoroute/1
    curl -X PUT http://localhost:80/${SWITCHER}/videoroute/1 -d "3"
    curl -X PUT http://localhost:80/${SWITCHER}/avroute/1 -d "2"
    curl http://localhost:80/${SWITCHER}/avroute/1
    curl http://localhost:80/${SWITCHER}/avroute/1
    curl http://localhost:80/${SWITCHER}/videoroute/1
    curl -X PUT http://localhost:80/${SWITCHER}/videoroute/8 -d "5"
    curl -X PUT http://localhost:80/${SWITCHER}/videoroute/8 -d "6"
    curl http://localhost:80/${SWITCHER}/videoroute/8
    curl http://localhost:80/${SWITCHER}/avroute/8
    curl http://localhost:80/${SWITCHER}/videoroute/8
    curl -X PUT http://localhost:80/${SWITCHER}/avroute/5 -d "8"
    curl -X PUT http://localhost:80/${SWITCHER}/avroute/5 -d "8"
    curl http://localhost:80/${SWITCHER}/videoroute/5
    curl http://localhost:80/${SWITCHER}/avroute/5
    curl http://localhost:80/${SWITCHER}/videoroute/5
    curl -X PUT http://localhost:80/${SWITCHER}/videoroute/7 -d "2"
    curl -X PUT http://localhost:80/${SWITCHER}/videoroute/7 -d "6"
    curl http://localhost:80/${SWITCHER}/videoroute/7
    curl http://localhost:80/${SWITCHER}/avroute/7
    curl http://localhost:80/${SWITCHER}/videoroute/7
    curl -X PUT http://localhost:80/${SWITCHER}/videoroute/2 -d "6"
    curl -X PUT http://localhost:80/${SWITCHER}/avroute/2 -d "5"
    curl http://localhost:80/${SWITCHER}/avroute/2
    curl http://localhost:80/${SWITCHER}/avroute/2
    curl http://localhost:80/${SWITCHER}/videoroute/6
    curl -X PUT http://localhost:80/${SWITCHER}/videoroute/6 -d "1"
    curl -X PUT http://localhost:80/${SWITCHER}/videoroute/6 -d "1"
    curl http://localhost:80/${SWITCHER}/videoroute/6
    curl http://localhost:80/${SWITCHER}/avroute/6
    curl http://localhost:80/${SWITCHER}/videoroute/7
    curl http://localhost:80/${SWITCHER}/avroute/7

    sleep 2
done
```
