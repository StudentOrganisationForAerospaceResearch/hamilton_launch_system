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

#define FIRE_DURATION 1000

#define RELAY_ON LOW
#define RELAY_OFF HIGH


SoftwareSerial umbilical(SERIAL_RX, SERIAL_TX);

bool armed = false;
bool fired  = false;
long int time;



//This function impliments the arm command
//The ARM relay is opened
void arm(){
  digitalWrite(ARM_RELAY, RELAY_ON);
  armed = true;
  umbilical.read(); //Read the remaining byte of the command
}

//This function impliments the fire command
//The FIRE relay is opened only if the system is armed
//otherwise the system will ignore this call
void fire(){
  if(armed){
    digitalWrite(FIRE_RELAY, RELAY_ON);
    fired = true;
    delay(FIRE_DURATION); //Wait for the igniter to get hot
    abort(); //Close the relays
  }
  umbilical.read(); //Read the remaining byte of the command
}

//This command will close both Arm and Fire relays
//In addition it will reset the sytem to a disarmed state
void abort(){
  digitalWrite(FIRE_RELAY, RELAY_OFF);
  digitalWrite(ARM_RELAY, RELAY_OFF);
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
  
  //Ensure the relays are closed
  digitalWrite(RELAY_1, RELAY_OFF);
  digitalWrite(RELAY_2, RELAY_OFF);
  digitalWrite(RELAY_3, RELAY_OFF);
  digitalWrite(RELAY_4, RELAY_OFF);
  digitalWrite(RELAY_5, RELAY_OFF);
  digitalWrite(RELAY_6, RELAY_OFF);   

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

  time = 0;
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
    byte header = Serial.read();
    if(header == 0x21) arm();
    else if(header == 0x20) fire();
    else if(header == 0x2F) abort();
    else while(Serial.available()>0){
      byte data = Serial.read();
      umbilical.write(data);
    }
  }
  
  if(millis()-time>1000){
    digitalWrite(LED,0x1^digitalRead(LED));
    time = millis();
  }
  
}
