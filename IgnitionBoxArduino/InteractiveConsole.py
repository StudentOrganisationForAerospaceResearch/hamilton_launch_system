import serial
import time
import binascii

order = 'little'


class AvionicsData:
    def __init__(self):
        self.imu = [0] * 9
        self.bar = [0] * 2
        self.gps = [0] * 4
        self.oxi = 0
        self.cmb = 0
        self.phs = -1
        self.vnt = 0

    def __str__(self):
        phases = ["NA!", "PRELAUNCH", "BURN", "COAST", "DROGUE_DESCENT", "MAIN_DESCENT", "ABORT"]

        string="IMU - ACCEL:\t"+ str(self.imu[0:3])+"\n"
        string+="IMU - GYRO:\t"+ str(self.imu[3:6])+"\n"
        string+="IMU - MAG:\t" + str(self.imu[6:9]) +"\n"
        string+="BAR - PRESS:\t"+ str(self.bar[0]) +"\n"
        string+="BAR - TEMP:\t"+  str(self.bar[1])+"\n"
        string+="GPS - ALT:\t"+ str(self.gps[0])+"\n"
        string+="GPS - TIME:\t"+ str(self.gps[1])+"\n"
        string+="GPS - LAT:\t"+ str(self.gps[2])+"\n"
        string+="GPS - LONG:\t"+ str(self.gps[3])+"\n"
        string+="OXI - PRESS:\t"+ str(self.oxi)+"\n"
        string+="CMB - PRESS:\t"+ str(self.cmb)+"\n"
        string+="VNT - STATUS:\t"+ str(self.vnt)+"\n"
        string+="PHS -  PHASE:\t"+ str(phases[self.phs + 1])+"\n"

        return string

def arm(ser):
    ser.write(bytearray.fromhex('2100'))
    time.sleep(1)
    print('Arm Command Sent!\n')

def fire(ser):
    ser.write(bytearray.fromhex('2000'))
    time.sleep(1)
    print('Fire Command Sent!\n')

def abort(ser):
    ser.write(bytearray.fromhex('2F00'))
    time.sleep(1)
    print('Abort Command Sent!\n')

def setBaud(ser,baud):
	ser.baudrate = baud
	time.sleep(1)
	ser.flush()
	print("Baudrate set to: ", baud)

def fillOpen(ser):
    ser.write(bytearray.fromhex('2200'))
    time.sleep(1)
    print('Fill Valve Open Sent!\n')

def fillClose(ser):
    ser.write(bytearray.fromhex('2300'))
    time.sleep(1)
    print('Fill Valve Close Sent!\n')


def help():
    print('\nList of commands:\n----------------------------------')
    print('abort\t\t send the abort command 0x2F')
    print('arm\t\t send the arm command 0x21')
    print('clear\t\t clears the terminal')
    print('disconnect\t disconnect and connect to another comm port')
    print('fire\t\t send the fire command 0x20')
    print('help\t\t prints this help menu')
    print('fill [open|close]\t opens or closes the Nitrous Fill Valve (command 0x22/0x23)')
    print('quit\t\t closes the serial terminal and program')
    print('read\t\t reads the serial buffer and displays the latest data\n')

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


def readHex(ser):
    line = ser.readline()
    print(binascii.hexlify(line))

def readSerial(ser,data):
    line = ser.readline()
    i = 0
    while(i<len(line)):
        if((line[i]==0x31) and (len(line)-i>=38)):
            for j in range(9):
                data.imu[j] = int.from_bytes(line[i+1+(j*4):line[i+5+(j*4)]], byteorder=order, signed=True)
            i+=38
        elif((line[i]==0x32) and (len(line)-i>=10)):
            data.bar[0] = int.from_bytes(line[i+1:i+5], byteorder=order, signed=True)
            data.bar[1] = int.from_bytes(line[i+5:i+9], byteorder=order, signed=True)
            i+=10
        elif((line[i]==0x33) and (len(line)-i>=18)):
            for j in range(9):
                data.imu[j] = int.from_bytes(line[i+1+(j*4):i+5+(j*4)], byteorder=order, signed=True)
            i+=18
        elif((line[i]==0x34) and (len(line)-i>=6)):
            data.oxi = int.from_bytes(line[i+1:i+4], byteorder=order, signed=True)
            i+=6
        elif((line[i]==0x35) and (len(line)-i>=6)):
            data.cmb = int.from_bytes(line[i+1:i+4], byteorder=order, signed=True)
            i+=6
        elif((line[i]==0x36) and (len(line)-i>=3)):
            data.phs = line[i+1]
            i+=3
        elif((line[i]==0x37) and (len(line)-i>=3)):
            data.vnt = line[i+1]
            i+=3
        else: i+=1

    print(data)

ser = None
data = AvionicsData()
while(True):

    while(ser!=None):
        comm = input("Awaiting command (enter help for list of commands):")
        if(comm == 'arm'): arm(ser)
        elif(comm == 'fire'): fire(ser)
        elif(comm == 'abort'): abort(ser)
        elif(comm == 'help'): help()
        elif(comm == 'disconnect'): ser = disconnect(ser)
        elif(comm == 'quit'): exit()
        elif(comm == 'hex'): readHex(ser)
        elif(comm == 'read'): readSerial(ser, data)
        elif(comm[0:4] == 'baud'): setBaud(ser, int(comm[5:]))
        elif(comm == 'fill open'): fillOpen(ser)
        elif(comm == 'fill close'): fillClose(ser)
        elif(comm == 'clear' or comm == 'cls'): print(chr(27) + "[2J")
        else: print(comm,': Command Not Found')

    ser = connect(input('Enter a Serial Port to connect to:'))
    help()
