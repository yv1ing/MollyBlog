#/bin/bash

kill `ps aux | grep 'molly' | awk '{print $2}' | head -n 1`