# AutoPunch - 打卡程式
## 開發環境
- Windows 10 x64 1909
- go1.15.5 windows/amd64

## 資料夾結構
```
|-- bin               #存放執行檔
|    |--punchin.exe   #簽到可執行檔
|    |--punchout.exe  #簽退可執行檔
|    |--info.json     #帳號密碼設定檔
|
|-- punch             #存放簽到簽退程式
|    |--punch.go      #簽到簽退程式
|
|-- punchin           #存放簽到主程式
|    |--main.go       #簽到主程式
|
|-- punchout          #存放簽退主程式
|    |--main.go       #簽退主程式
```
## 下載

### 整包下載
點選畫面右上角的`code`>`Download ZIP`即可下載所有檔案
![0](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/0.png?raw=true)

### 只下載必要的檔案
- 可執行檔
    - [punchin.exe](https://github.com/vincent87720/AutoPunch/raw/master/bin/punchin.exe)(簽到)
    - [punchout.exe](https://github.com/vincent87720/AutoPunch/raw/master/bin/punchout.exe)(簽退)
    - [info.json](https://raw.githubusercontent.com/vincent87720/AutoPunch/master/bin/info.json)(帳號密碼設定檔)
- 工作排程器檔案
    - [簽到.xml](https://github.com/vincent87720/AutoPunch/raw/master/taskschd/%E7%B0%BD%E5%88%B0.xml)(簽到工作)
    - [簽退.xml](https://github.com/vincent87720/AutoPunch/raw/master/taskschd/%E7%B0%BD%E9%80%80.xml)(簽退工作)

## 設定

### 設定帳號密碼
將bin資料夾內的info.json依照json格式設定
由於使用json格式讀入帳號密碼，故可以設定多筆帳號密碼供程式使用
```json
[
  {
    "id":"這裡放學號",
    "pw":"這裡放密碼"
  },
  {
    "id":"",
    "pw":""
  }
]
```
*須注意info.json必須與punchin.exe和punchout.exe位於相同目錄*

### 匯入排程工作
按下`windows鍵+R`開啟**執行**方塊，並輸入taskschd.msc開啟工作排程器
![1](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/1.png?raw=true)

點選右方**匯入工作**選項，將**簽到.xml**及**簽退.xml**檔案匯入
![2](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/2.png?raw=true)

![3](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/3.png?raw=true)

### 調整觸發排程
點選**觸發程序**頁籤可以看到所有會觸發該工作的時間點，可點選新增或編輯現有的項目
![4](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/4.png?raw=true)

### 設定執行檔位置
點選**動作**頁籤，選擇**啟動程式**動作後點選編輯>瀏覽，接著選擇`punchin.exe`或`punchout.exe`即可觸發打卡程式
![5](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/5.png?raw=true)

![6](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/6.png?raw=true)

![7](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/7.png?raw=true)

