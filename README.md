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
|
|-- taskschd          #存放工作排程器檔案
|    |--簽到.xml　     #工作排程器簽到設定檔
|    |--簽退.xml　     #工作排程器簽退設定檔
```
## 資料夾設定

下載完成後解壓縮檔案到想要存放的目錄(這裡以直接放在C:\目錄下為例)
![1](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/1.png?raw=true)

更改資料夾名稱為AutoPunch加上學號，例如:AutoPunch-F0708090
![14](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/14.png?raw=true)

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
![2](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/2.png?raw=true)

點選右方**匯入工作**選項，將**簽到.xml**及**簽退.xml**檔案匯入
![3](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/3.png?raw=true)

![4](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/4.png?raw=true)

### 調整使用者設定
要建立工作之前，必須先將使用者調整成本機帳戶才可以順利匯入工作

點選一般>安全性選項>變更使用者或群組
![10](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/10.png?raw=true)

點選進階
![11](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/11.png?raw=true)

點選立即尋找後，在下方列表選取目前的使用者帳戶，按下確定鍵
![12](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/12.png?raw=true)

![13](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/13.png?raw=true)

### 調整觸發排程
點選**觸發程序**頁籤可以看到所有會觸發該工作的時間點，可點選新增或編輯現有的項目
![5](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/5.png?raw=true)

預設設定延遲工作最高可達5分鐘，即表示隨機在設定的時間開始往後五分鐘內執行，若要在確切時間執行請取消勾選此選項
![15](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/15.png?raw=true)

### 設定執行檔位置
#### 尋找開始位置
進入程式執行檔**bin**資料夾中，對punchin.exe點選右鍵>內容，將該檔案的**位置**複製起來
![6](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/6.png?raw=true)

![7](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/7.png?raw=true)

#### 設定工作排程器

回到工作排程器，點選**動作**頁籤，點選**編輯**按鈕，在開始位置欄位填入先前複製的bin的路徑即可
![8](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/8.png?raw=true)

![9](https://github.com/vincent87720/AutoPunch/blob/master/assets/img/9.png?raw=true)