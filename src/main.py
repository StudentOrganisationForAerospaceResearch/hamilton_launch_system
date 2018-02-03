##!/usr/bin/python3

import sys

from hamilton_launch_system import hamilton_launch_system

if __name__ == '__main__':
    port = int(sys.argv[1])
    hamilton_launch_system.run(host='0.0.0.0', port=port)
