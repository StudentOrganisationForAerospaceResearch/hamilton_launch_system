#include <SoftwareSerial.h>

#define RELAY_1    2
#define RELAY_2    3
#define RELAY_3    4
#define RELAY_4    5
#define RELAY_5    6
#define RELAY_6    7
#define SERIAL_TX  9
#define SERIAL_RX  10
#define LED        13

#define USB_SERIAL_BAUD 115200
#define UMB_SERIAL_BAUD 115200

#define ARM_RELAY RELAY_1
#define FIRE_RELAY RELAY_2

SoftwareSerial umbilical(SERIAL_RX, SERIAL_TX);

bool armed = false;
bool fired  = false;


void arm(){
  digitalWrite(ARM_RELAY, HIGH);
  armed = true;
  umbilical.read();
}

void fire(){
  if(armed){
    digitalWrite(FIRE_RELAY,HIGH);
    fired = true;
    umbilical.read();
  }
}

void abort(){
  digitalWrite(FIRE_RELAY,LOW);
  digitalWrite(ARM_RELAY,LOW);
  armed = false;
}


void setup() {
  
  //Set all of the relay controls to output
  pinMode(RELAY_1, OUTPUT);
  pinMode(RELAY_2, OUTPUT);
  pinMode(RELAY_3, OUTPUT);
  pinMode(RELAY_4, OUTPUT);
  pinMode(RELAY_5, OUTPUT);
  pinMode(RELAY_6, OUTPUT);

  //Set the Arduino Onboard LED to output
  pinMode(LED,OUTPUT);
  
  //Set Umbilical Serial PinMode
  pinMode(SERIAL_RX, OUTPUT);
  pinMode(SERIAL_TX, OUTPUT);
  
  //Initialize the USB Serial Connection
  Serial.begin(USB_SERIAL_BAUD);
  while(!Serial){;} //Wait for USB Serial to connect
  
  //Initialize the Umbilical Serial Connection
  umbilical.begin(UMB_SERIAL_BAUD);


  digitalWrite(RELAY_1, LOW);
  digitalWrite(RELAY_2, LOW);
  digitalWrite(RELAY_3, LOW);
  digitalWrite(RELAY_4, LOW);
  digitalWrite(RELAY_5, LOW);
  digitalWrite(RELAY_6, LOW);   
}


void loop() {
  umbilical.listen();
  
  if(umbilical.available()>0){
    while(umbilical.available()>0){
      byte data = umbilical.read();
      Serial.write(data);
    }
  }
  
  if(Serial.available()>0){
    byte header = umbilical.read();
    if(header == 0x21) arm();
    else if(header == 0x20) fire();
    else if(header == 0x2F) abort();
    else while(Serial.available()>0){
      byte data = Serial.read();
      umbilical.write(data);
    }
  }
  
  
}
