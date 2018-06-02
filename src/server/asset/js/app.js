//source = new EventSource('/events');

function AlertBox(value){
    var alertBox = '<div class="alert-box error"><span></span>'+value+'</div>';
    return alertBox;
}

//source.addEventListener('update',function(e){
//    var data = JSON.parse(e.data);
    //console.log(data.Name);
//    clear();
//    AwesomeApp.$root.$children[0].add(data.Name);
//    console.log(AwesomeApp.$root.$children);
//    load(data.Name);
//    //d1.insertAdjacentHTML('beforeend',AlertBox(data));
//});

function ActivateAutorotate(){
    but = document.getElementById("autorotate");
    if (controls.autoRotate == true){
        but.classList.remove("button-success");
        controls.autoRotate = false;
    } else {
        but.classList.add("button-success");
        controls.autoRotate = true;
    }
}


Vue.config.productionTip = false;

AwesomeApp = new Vue({
    el: '#main',
	data: {
        modelList : [],
    },
    created() {
        this.setupStream();
    },
    methods: {
        setupStream() {
            // Not a real URL, just using for demo purposes
            let es = new EventSource('/events');
            es.onerror = function(e){
                console.log(e);
            }
            es.addEventListener('update', event => {
                let data = JSON.parse(event.data);
                this.modelList.push(data.Name);;
                clear();
                load(data.Name);
                console.log(data);
            }, false);

            es.addEventListener('error', event => {
                if (event.readyState == EventSource.CLOSED) {
                    console.log('Event was closed');
                    console.log(EventSource);
                }
            }, false);
        }
    }
});
