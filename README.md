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



### Copyright CC 2020 Walid Amriou

<a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by-nc-sa/4.0/88x31.png" /></a><br />This work is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/4.0/">Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License</a>.

__You are free to:__
  * Share — copy and redistribute the material in any medium or format
  * Adapt — remix, transform, and build upon the material
The licensor cannot revoke these freedoms as long as you follow the license terms.  

__Under the following terms:__
  * Attribution — You must give appropriate credit, provide a link to the license, and indicate if changes were made. You may do so in any reasonable manner, but not in any way that suggests the licensor endorses you or your use.
  * NonCommercial — You may not use the material for commercial purposes.
  * ShareAlike — If you remix, transform, or build upon the material, you must distribute your contributions under the same license as the original.
  * No additional restrictions — You may not apply legal terms or technological measures that legally restrict others from doing anything the license permits.


the software or Code or Scripts or any files in this source is provided "as is" and the author disclaims all warranties with regard to this files including all implied warranties of merchantability and fitness. in no event shall the author be liable for any special, direct, indirect, or consequential damages or any damages whatsoever resulting from loss of use, data or profits, whether in an action of contract, negligence or other tortious action, arising out of or in connection with the use or performance of this software or code or scripts or any files in this source.



