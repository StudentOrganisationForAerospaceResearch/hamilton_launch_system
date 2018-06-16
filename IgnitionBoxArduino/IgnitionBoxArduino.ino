#include <SoftwareSerial.h>
//#include <Wire.h>
#include <Adafruit_ADS1015.h>


#define RELAY_1    2
#define RELAY_2    3
#define RELAY_3    4
#define RELAY_4    5
#define RELAY_5    6
#define RELAY_6    7
#define SERIAL_TX  11
#define SERIAL_RX  10
#define AVIONICS_RESET 12
#define LED        13

#define USB_SERIAL_BAUD 9600
#define UMB_SERIAL_BAUD 9600

#define ARM_RELAY RELAY_1
#define FIRE_RELAY RELAY_2
#define FILL_RELAY RELAY_4
#define PURGE_RELAY RELAY_3
#define EJECT_RELAY RELAY_6

#define FIRE_DURATION 10000
#define PURGE_TIME 3000
#define EJECT_TIME 1000

#define RELAY_ON LOW
#define RELAY_OFF HIGH


SoftwareSerial umbilical(SERIAL_RX, SERIAL_TX);
Adafruit_ADS1115 ads;

bool armed = false;
bool fired  = false;
long int time;

//This function impliments the arm command
//The ARM relay is opened
void arm(){
  digitalWrite(FILL_RELAY, RELAY_OFF);
  digitalWrite(PURGE_RELAY, RELAY_ON);
  delay(PURGE_TIME);
  digitalWrite(PURGE_RELAY, RELAY_OFF);
  digitalWrite(EJECT_RELAY, RELAY_ON);
  delay(EJECT_TIME);
  digitalWrite(EJECT_RELAY, RELAY_OFF);
  digitalWrite(ARM_RELAY, RELAY_ON);
  armed = true;
  Serial.read(); //Read the remaining byte of the command
}

//This function impliments the fire command
//The FIRE relay is opened only if the system is armed
//otherwise the system will ignore this call
void fire(){
  if(armed){
    umbilical.write((byte) 0x20);
    umbilical.write((byte) 0x00);
    digitalWrite(FIRE_RELAY, RELAY_ON);
    fired = true;
    delay(FIRE_DURATION); //Wait for the igniter to get hot
    abortCommand(); //Close the relays
  }
  Serial.read(); //Read the remaining byte of the command
}

//This command will close both Arm and Fire relays
//In addition it will reset the sytem to a disarmed state
void abortCommand(){
  umbilical.write((byte) 0x2F);
  umbilical.write((byte) 0x00);
  digitalWrite(FIRE_RELAY, RELAY_OFF);
  digitalWrite(ARM_RELAY, RELAY_OFF);
  armed = false;
  Serial.read();
}

void openFillValve(){
  digitalWrite(FILL_RELAY, RELAY_ON);
  Serial.read();  
}

void closeFillValve(){
  digitalWrite(FILL_RELAY, RELAY_OFF);
  Serial.read();
}

void setup() {
  
  //Set all of the relay controls to output
  pinMode(RELAY_1, OUTPUT);
  pinMode(RELAY_2, OUTPUT);
  pinMode(RELAY_3, OUTPUT);
  pinMode(RELAY_4, OUTPUT);
  pinMode(RELAY_5, OUTPUT);
  pinMode(RELAY_6, OUTPUT);

  pinMode(AVIONICS_RESET, INPUT_PULLUP);
  
  //Ensure the relays are closed
  digitalWrite(RELAY_1, RELAY_OFF);
  digitalWrite(RELAY_2, RELAY_OFF);
  digitalWrite(RELAY_3, RELAY_OFF);
  digitalWrite(RELAY_4, RELAY_OFF);
  digitalWrite(RELAY_5, RELAY_OFF);
  digitalWrite(RELAY_6, RELAY_OFF);   

  //Set the Arduino Onboard LED to output
  pinMode(LED,OUTPUT);
  
  //Initialize the USB Serial Connection
  Serial.begin(USB_SERIAL_BAUD);
  while(!Serial){
    ;
  } //Wait for USB Serial to connect
  
  //Initialize the Umbilical Serial Connection
  umbilical.begin(UMB_SERIAL_BAUD);
  umbilical.write((byte) 0);

  time = 0;
}


void loop() {
  
  if(umbilical.available()){
      Serial.write(umbilical.read());
  }
  
  if(Serial.available()){
    byte header = Serial.read();
    if(header == 0x21) arm();
    else if(header == 0x20) fire();
    else if(header == 0x2F) abortCommand();
    else if(header == 0x22) openFillValve();
    else if(header == 0x23) closeFillValve();
    else umbilical.write(header);
  }

  if(digitalRead(AVIONICS_RESET)==LOW){
    abortCommand();
    delay(100);
    umbilical.write((byte) 0x45);
    umbilical.write((byte) 0x00);  
  }

  
  if(millis()-time>1000){
    if(analogRead(0)>350){  //If the voltage is above nominal levels
      //Send a heartbeat packet
      umbilical.write((byte) 0x46);
      umbilical.write((byte) 0x00);
      digitalWrite(LED,0x1^digitalRead(LED));      
    }
    else{
      digitalWrite(LED,HIGH);
    }
    time = millis();

  }

  
  
  
  
}
