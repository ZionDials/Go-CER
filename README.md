# Go-CER

Go-CER creates an idle application for Cisco IP phones that displays the current status of the CER database. It is designed to be used with the [Cisco Emergency Responder](https://www.cisco.com/c/en/us/products/unified-communications/emergency-responder/index.html) product.

## Usage

``` bash
go-cer serve --config "dev.yaml"
```

## Limitations

- Cisco IP Phones only query the XML application on idle. This means that the status will only be updated when the phone goes to idle again.

## Failed Response

``` xml
<CiscoIPPhoneImageFile>
   <LocationX>0</LocationX>
   <LocationY>0</LocationY>
   <Width>800</Width>
   <Height>480</Height>
   <Title />
   <Prompt>Not Provisioned for CER</Prompt>
   <URL>website.failureUrl</URL>
   <SoftKeyItem>
      <Name>Exit</Name>
      <URL>SoftKey:Exit</URL>
      <Position>2</Position>
   </SoftKeyItem>
</CiscoIPPhoneImageFile>
```

## Success Response

``` xml
<CiscoIPPhoneImageFile>
   <LocationX>0</LocationX>
   <LocationY>0</LocationY>
   <Width>800</Width>
   <Height>480</Height>
   <Title />
   <Prompt>ERL-Name</Prompt>
   <URL>website.sucessUrl</URL>
   <SoftKeyItem>
      <Name>Exit</Name>
      <URL>SoftKey:Exit</URL>
      <Position>2</Position>
   </SoftKeyItem>
</CiscoIPPhoneImageFile>
```

## Example Config

``` yaml
cer:
  username: ceridle # Username for the CER website
  password: 012345abc # Password for the CER website
  host: localhost # Host of the CER website
database:
  autoMigrate: true # Migrate the database schema on startup
  database: cer # Database name
  driver: postgres # Database driver (mysql|mssql|postgres|sqlite)
  host: localhost # Database host
  limit: 100 # Maximum number of records to insert in bulk
  password: 012345abc # Database password
  path: ./go-cer.db # Path to the database file (sqlite only)
  port: 5432 # Database port
  username: postgres # Database username
  SSL: disable # Database SSL mode (disable|require|verify-ca|verify-full)
logging:
  compress: true # Compress log files
  level: info # Logging level (debug|info|warn|error|fatal|panic)
  maxAge: 30 # Maximum age of log files in days
  maxSize: 100 # Maximum size of log files in megabytes
  name: go-cer.log # Name of the log files
  path: ./logs # Path to store log files
website:
  port: 8080 # Port to listen on
  title: Go-CER # Title of the website
  sucessUrl: http://localhost:8080 # URL of the website image for success
  failureUrl: http://localhost:8080 # URL of the website image for failure
  softkeys:
   - Name: Exit
     URL: SoftKey:Exit
     Position: 2
```
