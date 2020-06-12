# Moskotcho_Slides_Remote
Remote control slides presentations From Web browsers &amp; any device support WiFi 

### To develop
```bash
# 1- You need to have Golang install in your local machine.    
https://lokumgolang.walidamriou.com/Install_Go_language.html   

# 2- Clone the project  
git clone https://github.com/walidamriou/Moskotcho_Slides_Remote  

# 3- Install packages:  
go get github.com/bmizerany/pa  
# then  
go get github.com/go-vgo/robotgo  

# 4- open the terminal in Server folder and write:   
go run main.go    

# 5- Open your slide in presentation mode (I tesed it with PDf file and ubuntu 20.04 Document Viewer)  

# 5- open the browser in any other device that support WiFi and browsers and write http://ip_of_you_local_machine:8000 

# 6- and you will get a web page like that:   
```
<p align="center">
  <img width="50%" height="50%" src="https://github.com/walidamriou/Moskotcho_Slides_Remote/blob/master/img/localhost_8000_(Galaxy%20S5).png">
</p>
```
# 7- Test click next/back buttons   
```

:warning: If there are any issues in the scripts of the examples write an issue in: [issues](https://github.com/walidamriou/LokumGoLang/issues "issues")   

### You want to integrete this project with other project ?
If you want to integerete this server in other projects like android application or IoT device based Raspberry pi or ESP32 ...
you just make a http post request with the URL: http://ip_of_you_local_machine:8000 and in body send: 
for Next button "button=next" and for back button send "button=back".

### Website of the project:  
https://moskotchosr.walidamriou.com/

<p align="center">
  <img width="50%" height="50%" src="https://github.com/walidamriou/Moskotcho_Slides_Remote/blob/master/img/anim.png">
</p>

### If you need any help or informations:
:large_blue_circle:	 Facebook: https://www.facebook.com/walidamriou   
:large_blue_circle:  Twitter: https://twitter.com/walidamriou    
:red_circle: Email:  contact [at] walidamriou [dot] com    






