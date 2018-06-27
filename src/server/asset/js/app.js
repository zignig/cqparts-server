
function AlertBox(value){
    var alertBox = '<div class="alert-box error"><span></span>'+value+'</div>';
    return alertBox;
}

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
// shh
Vue.config.productionTip = false;
// add semantic ui
Vue.use(SemanticUIVue);

// create and event bus
EventBus = new Vue();

AwesomeApp = new Vue({
    el: '#main',
	data: {
        modelList : [],
        issueItem : '',
    },
    created() {
        this.setupStream();
    },
    methods: {
        setupStream() {
            let es = new EventSource('/events');
                es.onerror = function(e){
                console.log(e);
            }

            es.addEventListener('menu', event => {
                let data = JSON.parse(event.data);
                if ((data.Name) != ''){
                    this.modelList.push(data.Name);
                };
            }, false);

            es.addEventListener('update', event => {
                let data = JSON.parse(event.data);
                len = this.modelList.length
                EventBus.$emit('menu item',event.data);
                if ((data.Name) != ''){
                    if (this.modelList.includes(data.Name) == false){
                        this.modelList.unshift(data.Name);
                    } else {
                        console.log("fnord");
                        if (this.modelList.indexOf(data.Name) > 0) {
                            this.modelList.splice(this.modelList.indexOf(data.Name), 1);
                            this.modelList.unshift(data.Name);
                        }
                    }
                    if (this.modelList.length > 15){
                        this.modelList.slice(0,15)
                    }
                };
                this.issueItem = '';
                clear();
                load(data.Name);
            }, false);

            es.addEventListener('issue', event => {
                let data = JSON.parse(event.data);
                this.issueItem = data.Name;
                EventBus.$emit('issue',data.Name);
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
