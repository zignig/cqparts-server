function sendBinary(data, url){
	var xhr = new XMLHttpRequest();
	xhr.open("POST", url, true);

	if (typeof XMLHttpRequest.prototype.sendAsBinary == "function") { // Firefox 3 & 4
		var tmp = '';
		for (var i = 0; i < data.length; i++) tmp += String.fromCharCode(data.charCodeAt(i) & 0xff);
		data = tmp;
	}
	else { // Chrome 9
		// http://javascript0.org/wiki/Portable_sendAsBinary
		XMLHttpRequest.prototype.sendAsBinary = function(text){
			this.send(data);
		}
	}
	xhr.sendAsBinary(data);	
}

function Snapshot(){
	var strMime = "image/jpeg";
	imgData = renderer.domElement.toDataURL(strMime);
	sendBinary(imgData, "snapshot");
}
