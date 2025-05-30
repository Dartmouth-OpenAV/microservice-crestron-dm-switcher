#!/bin/bash

# Crestron DM Switcher microservice test script
# Replace these variables with your actual values
MICROSERVICE_URL="your-microservice-url"
DEVICE_FQDN="your-device-fqdn"

echo "Testing Crestron DM Switcher microservice..."
echo "Microservice URL: $MICROSERVICE_URL"
echo "Device FQDN: $DEVICE_FQDN"
echo "----------------------------------------"

# GET avroute
echo "Testing GET avroute..."
curl -X GET "http://${MICROSERVICE_URL}/${DEVICE_FQDN}/avroute/1"
sleep 1

# SET avroute
echo "Testing SET avroute..."
curl -X PUT "http://${MICROSERVICE_URL}/${DEVICE_FQDN}/avroute/1" \
     -H "Content-Type: application/json" \
     -d '"2"'
sleep 1

# GET videoroute
echo "Testing GET videoroute..."
curl -X GET "http://${MICROSERVICE_URL}/${DEVICE_FQDN}/videoroute/1"
sleep 1

# SET videoroute
echo "Testing SET videoroute..."
curl -X PUT "http://${MICROSERVICE_URL}/${DEVICE_FQDN}/videoroute/1" \
     -H "Content-Type: application/json" \
     -d '"2"'
sleep 1

echo "----------------------------------------"
echo "All tests completed."
