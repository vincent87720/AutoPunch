package punch

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"regexp"
	time "time"
)

func Punch(username string, password string, punchType string) {
	startDate, endDate := getDateRange()
	
	//登入請求的目標URL
	loginURL := "http://ais.dyu.edu.tw/prj_epfee/login_lib.php" 
	//簽到簽退的目標URL
	punchURL := "http://ais.dyu.edu.tw/prj_epfee/prj_carddata_list.php?func_id=frm_ins&txt_bdate=" + startDate + "&txt_edate=" + endDate + "&txt_stno=" + username + "&hdn_itype=" + punchType
	//登入時的驗證表單資訊
	validateFields := url.Values{"txt_userid": {username}, "pwd_password": {password}}
	//簽到簽退的驗證表單資訊
	punchFields := url.Values{"txt_bdate": {startDate}, "txt_edate": {endDate}, "txt_stno": {username}, "func_id": {"frm_chktime"}, "hdn_itype": {punchType}}

	//發送登入請求
	response, err := http.PostForm(loginURL, validateFields)
	if err != nil {
		log.Fatal(err)
	}
	//暫存cookies供後續請求使用
	savedCookies := response.Cookies()

	//因http.NewRequest方法最後一個參數需傳入io.Reader類型的body內容
	//所以將簽到簽退的驗證表單資訊Encode後經由NewReader轉成Reader類型並放入body
	body := strings.NewReader(punchFields.Encode())
	//建立簽到簽退請求
	request, err := http.NewRequest("POST", punchURL, body)
	if err != nil {
		log.Fatal(err)
	}
	//將先前儲存在savedCookies的cookies加入到請求中
	for _, cookie := range savedCookies {
		request.AddCookie(cookie)
	}
	client := &http.Client{}
	//發送簽到簽退請求
	response, err = client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status: "+response.Status)
	// respBody, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// checkPunchStatus(respBody)
}

func getDateRange() (startDate string, endDate string) {
	timeStr := time.Now().String()
	monthNumStr := timeStr[5:7]
	year, month, _ := time.Now().Date()
	leapYear := judgeLeapYear(year)
	if leapYear && month == time.February {
		startDate = strconv.Itoa(year) + "/" + monthNumStr + "/01"
		endDate = strconv.Itoa(year) + "/" + monthNumStr + "/29"
	} else {
		switch month {
		case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
			startDate = strconv.Itoa(year) + "/" + monthNumStr + "/01"
			endDate = strconv.Itoa(year) + "/" + monthNumStr + "/31"
			break
		case time.April, time.June, time.September, time.November:
			startDate = strconv.Itoa(year) + "/" + monthNumStr + "/01"
			endDate = strconv.Itoa(year) + "/" + monthNumStr + "/30"
			break
		case time.February:
			startDate = strconv.Itoa(year) + "/" + monthNumStr + "/01"
			endDate = strconv.Itoa(year) + "/" + monthNumStr + "/28"
			break
		}
	}
	return
}

func judgeLeapYear(year int) (leapYear bool) {
	leapYear = false //1為閏年,0為平年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		leapYear = true
	}
	return leapYear
}

func checkPunchStatus(body []byte){
	reg,err := regexp.Compile(`alert\('簽退成功'\)`)
	if err != nil {
		log.Fatal(err)
	}
	regexpPnhInSuccess,err := regexp.Compile(`&hdn_itype=A`)
	if err != nil {
		log.Fatal(err)
	}
	if reg.Match(body){
		fmt.Println("Punch: 簽退成功")
	}else if regexpPnhInSuccess.Match(body){
		fmt.Println("Punch: 簽到成功")
	}else{
		fmt.Println("Punch: 簽到/簽退失敗")
	}
}