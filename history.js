
//Scrape history from for bot https://www.playok.com/en/turkishdama
function getHist(){
	let historyTable = document.getElementsByClassName("br")[0].children[0]

	var historyText = ""
	for(item of historyTable.children){
		if(item.children[2] != undefined) historyText += item.children[1].innerText + " " + item.children[2].innerText + " "
		else historyText += item.children[1].innerText
	}
	return historyText.trim()
}

console.log(getHist())