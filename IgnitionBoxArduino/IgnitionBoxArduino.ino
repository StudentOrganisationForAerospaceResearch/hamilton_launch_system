#include <SoftwareSerial.h>
#include <Wire.h>
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


#define HEART_BEAT_DELAY 5000
#define LOAD_CELL_TIME_DELAY 500

#define FIRE_DURATION 10000
#define PURGE_TIME 3000
#define EJECT_TIME 1000


//Load Cell constants
#define LOAD_CELL_A 1
#define LOAD_CELL_B 0

#define LOAD_CELL_BUFFER_SIZE 10

#define RELAY_ON LOW
#define RELAY_OFF HIGH


SoftwareSerial umbilical(SERIAL_RX, SERIAL_TX);
Adafruit_ADS1115 ads;

bool armed = false;
bool fired  = false;
long int heartTime;
long int loadTime;

int32_t loadBuffer[LOAD_CELL_BUFFER_SIZE];
long int loadTotal;
int loadPtr;


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

void readLoadCell(){
    //Calculate the load Cell packet
    loadTotal-=loadBuffer[loadPtr];
    loadBuffer[loadPtr] = ads.readADC_Differential_0_1();
    loadTotal+=loadBuffer[loadPtr];
    loadPtr = (loadPtr+1) % LOAD_CELL_BUFFER_SIZE;
    
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

  ads.setGain(GAIN_ONE);
  ads.begin();
  
  //Initialize the USB Serial Connection
  Serial.begin(USB_SERIAL_BAUD);
  while(!Serial){
    ;
  } //Wait for USB Serial to connect
  
  //Initialize the Umbilical Serial Connection
  umbilical.begin(UMB_SERIAL_BAUD);
  umbilical.write((byte) 0);

  heartTime = 0;
  loadTime = 0;

  for(int i=0; i<LOAD_CELL_BUFFER_SIZE; i++) loadBuffer[i]=0;
  loadPtr=0;
  loadTotal=0;
}


void loop() {

  
  //Feed all commands from avionics to the ground station
  if(umbilical.available()){
    while(umbilical.available()){
      Serial.write(umbilical.read());
    }
  }
  


  //Recive and parse commands from ground station
  if(Serial.available()){
    while(Serial.available()){
      byte header = Serial.read();
      if(header == 0x21) arm();
      else if(header == 0x20) fire();
      else if(header == 0x2F) abortCommand();
      else if(header == 0x22) openFillValve();
      else if(header == 0x23) closeFillValve();
      else umbilical.write(header);
    }
  }

  //Check for the Avionics reset
  if(digitalRead(AVIONICS_RESET)==LOW){
    delay(100); //Debounce delay
    if(digitalRead(AVIONICS_RESET)==LOW){
      abortCommand();
      delay(100);
      umbilical.write((byte) 0x45);
      umbilical.write((byte) 0x00); 
    }
  }

  //Send a Heartbeat packet
  if(millis()-heartTime >= HEART_BEAT_DELAY){
    if(analogRead(0)>350){  //If the voltage is above nominal levels
      readLoadCell();
      //Send a heartbeat packet
      umbilical.write((byte) 0x46);
      umbilical.write((byte) 0x00);
      digitalWrite(LED,0x1^digitalRead(LED));
    }
    else{
      digitalWrite(LED,HIGH);
    }
    heartTime = millis();
  }


  //Send a Load Cell packet
  if(millis()-loadTime>LOAD_CELL_TIME_DELAY){
    if(analogRead(0)>350){  //If the voltage is above nominal levels
      readLoadCell();
      
      int32_t loadCell = (int32_t) (loadTotal/LOAD_CELL_BUFFER_SIZE);
      
      //Send the loadcell packet
      for(int i=0; i<4; i++) Serial.write(0x40);
      Serial.write((loadCell>>24) && 0xff); 
      Serial.write((loadCell>>16) && 0xff);
      Serial.write((loadCell>>8) && 0xff);
      Serial.write(loadCell && 0xff);     
      Serial.write((byte) 0x00);
      

    }
    loadTime=millis();


  }

  
  
  
  
}
