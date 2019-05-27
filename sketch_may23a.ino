#include <ArduinoJson.h>
#include <SPI.h>
#include <MFRC522.h>
#include <ESP8266HTTPClient.h>
#include <ESP8266WiFi.h>
#define RST_PIN         5          // Configurable, see typical pin layout above
#define SS_PIN          15         // Configurable, see typical pin layout above

MFRC522 mfrc522(SS_PIN, RST_PIN);  // Create MFRC522 instance
String payload="";
String content="";
int hCode;
StaticJsonBuffer<200> jsonBuffer;
HTTPClient ati,cfe;
void setup() {
  Serial.begin(9600);   // Initialize serial communications with the PC
  WiFi.begin("Moto G (4) 8562","qwertybnm");
  while(WiFi.status()!=WL_CONNECTED){
    delay(500);
    Serial.println("Connecting");    
    }
    
  while (!Serial);    // Do nothing if no serial port is opened (added for Arduinos based on ATMEGA32U4)
  SPI.begin();      // Init SPI bus
  mfrc522.PCD_Init();   // Init MFRC522
  mfrc522.PCD_DumpVersionToSerial();  // Show details of PCD - MFRC522 Card Reader details
  Serial.println(F("Scan PICC to see UID, SAK, type, and data blocks..."));
}

void loop() {
  // Reset the loop if no new card present on the sensor/reader. This saves the entire process when idle.
  if ( ! mfrc522.PICC_IsNewCardPresent()) {
    return;
  }

  // Select one of the cards
  if ( ! mfrc522.PICC_ReadCardSerial()) {
    return;
  }
  
  for(byte i=0; i<mfrc522.uid.size; i++){
    content.concat(String(mfrc522.uid.uidByte[i],HEX));
    }
  content.toUpperCase();  
  Serial.println("UID-->"+content);//rfid_no
  delay(1000);
  
  
  while(content!=""){
  ati.begin("http://54.189.131.192:8192/inventory");
  ati.addHeader("Content-Type","application/json");
  JsonObject& root=jsonBuffer.createObject();
  root["rfid_no"]=content;
  root["barcode"]="ABC";
  String root1="";
  root.printTo(root1);
  //String root1="{\"rfid_no\":"+content+"}";
  Serial.println("Root 1----->"+root1);
   hCode= ati.POST(root1);
   content="";
   if(hCode>0){
   payload =ati.getString();
   Serial.println(payload);
   if (payload.equalsIgnoreCase("OK")){
   Serial.println("Successfully posted at ATI");
   Serial.println("hcode---->"+hCode);
   content="";
    }
   }    
    delay(2000);
    ati.end();
  
  if(payload.equalsIgnoreCase("Exist")){
  cfe.begin("http://54.189.131.192:8192/exit");
  cfe.addHeader("Content-Type","application/json"); 
  JsonObject& root_cfe=jsonBuffer.createObject();
  root_cfe["rfid_no"]=content;
  String root2="";
  root_cfe.printTo(root2);
  //String root2="{\"rfid_no\":"+content+"}";  
  Serial.println("Root 2----->"+root2);
   hCode= cfe.POST(root2);
   if(hCode>0){
   payload =cfe.getString();
   Serial.println(payload);
   if(payload.equalsIgnoreCase("OK_INV")){
    Serial.println("Posted at CFE");
    content="";   
    }
   }else{
    Serial.println("Failed at CFE object");
    Serial.println(hCode);
    content="";
    }
    delay(2000);
  cfe.end();    
   }
}
 delay(100);  
  }


