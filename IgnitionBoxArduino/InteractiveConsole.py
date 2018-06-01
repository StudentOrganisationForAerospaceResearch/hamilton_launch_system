import serial
import time

def arm(ser):
    ser.write(bytearray.fromhex('21'))
    time.sleep(1)
    print('\nArm Command Sent!')

def fire(ser):
    ser.write(bytearray.fromhex('20'))
    time.sleep(1)
    print('\nFire Command Sent!')

def abort(ser):
    ser.write(bytearray.fromhex('2F'))
    time.sleep(1)
    print('\nAbort Command Sent!\n')

def help():
    print('\nList of commands:\n----------------------------------')
    print('arm\t send the arm command 0x21')
    print('fire\t send the arm command 0x20')
    print('abort\t send the arm command 0x2F')
    print('help\t prints this help menu')
    print('disconnect\t disconnect and connect to another comm port')
    print('quit\t closes the serial terminal and program\n')


def connect(port):
    ser = serial.Serial(port, 115200, timeout=0)
    time.sleep(2)
    return ser

def disconnect(ser):
    ser.close()
    time.sleep(1)
    return None

def quit():
    exit()


ser = None
while(True):

    while(ser!=None):
        comm = input("Awaiting command (enter help for list of commands):")
        if(comm == 'arm'): arm(ser)
        elif(comm == 'fire'): fire(ser)
        elif(comm == 'abort'): abort(ser)
        elif(comm == 'help'): help()
        elif(comm == 'disconnect'): ser = disconnect(ser)
        elif(comm == 'quit'): exit()
        else: print(comm,': Command Not Found')

    ser = connect(input('Enter a Serial Port to connect to:'))