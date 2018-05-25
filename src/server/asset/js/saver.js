function sendBinary(data, url){
	var xhr = new XMLHttpRequest();
	xhr.open("POST", url, true);
    XMLHttpRequest.prototype.sendAsBinary = function(text){
        this.send(data);
    }
	xhr.sendAsBinary(data);	
}

function Snapshot(){
	var strMime = "image/jpeg";
	imgData = renderer.domElement.toDataURL(strMime);
	sendBinary(imgData, "snapshot");
}
