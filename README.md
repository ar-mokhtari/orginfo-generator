<!--suppress ALL -->
<h1 align="center" style="alignment: center;color: yellow">

orgInfo-generator
</h1>
<p align="center"><b>backend</b> (Golang)<br/><br/><br/></p>

<p align="center">
<a href="" target="_blank"><img src="https://img.shields.io/badge/Go-1.19+-00ADD8?style=for-the-badge&logo=go" alt="go version" /></a>&nbsp;
</p>

<h3 align="center">orginfo-generator is a generator tool, which create a back-end structure</h3>






<table cellpadding="0" cellspacing="0">
<tr>
    <td style="margin-top: -2rem" valign="button"><h4 style="color: deepskyblue;margin-bottom: -1.5rem">Powered by</h4></td>
    <td rowspan="2"><img alt="cgapp logo" src="https://raw.githubusercontent.com/create-go-app/cli/master/.github/images/cgapp_logo%402x.png" width="100px"/><br/></td>
  </tr>
  <tr>
    <td><img src="https://camo.githubusercontent.com/4724436344c2473558068577d7e9e6b597c2baabe75a499cd67e04a448e00d84/68747470733a2f2f7777772e766563746f726c6f676f2e7a6f6e652f6c6f676f732f676f6c616e672f676f6c616e672d617232312e737667" ></td>
  </tr>
</table> 



---

<h6> please notice me if you have any other idea ...
</h6>

<a href="mailto:neatland@gmail.com"><img alt="Gmail" title="Alireza Mokhtari G Gmail" src="https://camo.githubusercontent.com/571384769c09e0c66b45e39b5be70f68f552db3e2b2311bc2064f0d4a9f5983b/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f476d61696c2d4431343833363f7374796c653d666f722d7468652d6261646765266c6f676f3d676d61696c266c6f676f436f6c6f723d7768697465" data-canonical-src="https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&amp;logo=gmail&amp;logoColor=white" style="max-width: 100%;"></a>
<a href="https://t.me/ar_mokhtari" rel="nofollow"><img alt="Telegram" title="Alireza Mokhtari G Telegram" src="https://camo.githubusercontent.com/cf4ed981404024c1adfc79d5575c4edf1836c4fe36b24b03383ece888cef7e29/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f54656c656772616d2d3243413545303f7374796c653d666f722d7468652d6261646765266c6f676f3d74656c656772616d266c6f676f436f6c6f723d7768697465" data-canonical-src="https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&amp;logo=telegram&amp;logoColor=white" style="max-width: 100%;"></a>
<a href="https://www.linkedin.com/in/alireza-mokhtari-garakani-b4288024/"><img src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&amp;logo=linkedin&amp;logoColor=white" style="max-width: 100%;"></a>
<a href="https://discord.gg/F2YVf8Bu4R"><img src="![Discord](https://img.shields.io/badge/%3CServer%3E-%237289DA.svg?style=for-the-badge&logo=discord&logoColor=white" style="max-width: 100%;"></a>

---

<h3> Development map </h3>
<hr>

````
mkdir orginfo && cd orginfo
go mod init orginfo
------------------------------------------echo
go get github.com/labstack/echo/v4
------------------------------------------gorm
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
------------------------------------------ozzo-validation
go get github.com/go-ozzo/ozzo-validation

````

<h3>Domain generator</h3>
<hr>
to create new domain with CRUD operation and other prerequisite:

````
mkdir project_name
mkdir project_name/cli/generator 
touch main.go
````

import :

[domain generator](cli/generator/main.go)
and follow command(s) ...
like run this from '/config/cli/generator/bin/':

````
 cd go/src/orginfo/config/cli/generator/bin/
````

<h5>pattern #1</h5>
input from command and flag one by one:

````
./generator sub-command -domain_name="DOMAIN NAME" -fields="field1-string-field1_1,field2-uint-field1_2,..."
````

run this for add **a** new domain:

````
./generator new -domain_name="DOMAIN_NAME" -fields="codeType-uint-code_type,code-uint-code"
````

Also, to remove a domain run this:

````
./generator delete -domain_name="DOMAIN_NAME"
````

<h5>pattern #2</h5>
<h6>import domain(s) from json type file ... </h6>

````
./generator new -from_file
````

<h6>remove domain(s) from json type file ... </h6>

````
./generator delete -from_file
````

<h6>the file name must be: `input.json`. in this address: `/config/cli/generator/` </h6>
<h6>you can find default of file in this address: [sample structure](cli/generator/input_struct.json)</h6>

```json
[
  {
    "domain_name": "string",
    "fields": [
      {
        "name": "string",
        "type": "string",
        "json_name": "string"
      }
    ]
  }
]

```

https://user-images.githubusercontent.com/49469395/194825671-d93fe8e4-64fb-4c67-a8fb-c5e29e55e16c.mp4

<h3> Create and set service for production </h3>
<hr>
<h4> Create service </h4>

````
nano /lib/systemd/system/orginfo.service
````

<h4> Set service </h4>

````
[Unit]
Description=orginfo

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/home/aliz/go/src/orginfo/bin/orginfo

[Install]
WantedBy=multi-user.target
````

<h4> Service's actions </h4>

````
service orginfo start
service orginfo restart
service orginfo status
service orginfo enable
service orginfo stop
````

<h3> SDKs setting </h3>
<hr>
<ul></ul>

<h3> Environment </h3>
<ul></ul>

<h3>Json collection</h3>
To use and call from API platform like "Postman":

````
/docs/orginfo.postman_collection.json
````

<h3>Git</h3>
<hr>
<p align="left">
<a href="" 
target="_blank"><img src="https://img.shields.io/badge/git-%23F05033.svg?style=for-the-badge&logo=git&logoColor=white" alt="git" /></a>&nbsp;
</p>

for clone

````
 git clone git@github.com:ar-mokhtari/orginfo.git
````

for production, there is stable **main** branch:

```` 
 git checkout main
 
````

for develop:

```` 
 git checkout -b developing
 
````

after all, back to main:

```` 
 git merge --no-ff developing
 
````

<h3>Docker (compose)</h3>
<hr>

<a href="" target="_blank"><img src="https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white" /></a>
&nbsp;

for detail see:
[docker-compose.yaml](docker-compose.yaml)

````
docker compose up -d
````
