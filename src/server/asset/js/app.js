
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

function EventToServer(section,name){
    axios.post('/postevent',{
        section: section,
        name: name,
    });
}
function PostRender(){
    axios.post('/postrender',{
        name: vm.current,
        cam : camera.position,
        target : controls.target
    });
}

// shh
Vue.config.productionTip = false;
// add semantic ui
Vue.use(SemanticUIVue);
//Vue.use(Vuex);

// create and event bus
EventBus = new Vue();

vm = new Vue({
    el: '#main',
    store,
    data: {
        modelList : [],
        issueItem : '',
        current : '',
        modeltree : {'children':[]}, 
    },
    created() {
        this.setup();
    },
    methods: {
        setCurrent : function (obj) {
            this.current = obj;
        },
        setTree: function (obj) {
            this.modeltree = obj;
        },
        setup : function () {
            // base load
            axios.get('/list')
                .then( response => (this.modelList = response.data));

            EventBus.$on('pin',function(payload){
                EventToServer('pin',payload);
            });
            // select a model, set camera and dislpay
            EventBus.$on('select',function(payload){
                clear();
                camera.position.x = payload.view.cam.x;
                camera.position.y = payload.view.cam.y;
                camera.position.z = payload.view.cam.z;
                controls.target.x = payload.view.target.x;
                controls.target.y = payload.view.target.y;
                controls.target.z = payload.view.target.z;
                load(payload.name);
                vm.setCurrent(payload.name);
            });
            EventBus.$on('render',function(payload){
                vm.setCurrent(payload.name);
                PostRender();
            });

            let es = new EventSource('/events');
            es.onerror = function(e){
                console.log(e);
            }

            es.addEventListener('menu', event => {
                let data = JSON.parse(event.data);
                EventBus.$emit('menu item',data);
		console.log(data);
                if ((data.Name) != ''){
			clear();
			load(data.name);
                };
            }, false);

            es.addEventListener('update', event => {
                let data = JSON.parse(event.data);
                len = this.modelList.length
                EventBus.$emit('menu item',4);
                if ((data.Name) != ''){
                    if (this.modelList.includes(data) == false){
                        this.modelList.unshift(data);
                        this.current = data.name;
                    } else {
                        console.log("fnord");
                        if (this.modelList.indexOf(data) > 0) {
                            this.modelList.splice(this.modelList.indexOf(data), 1);
                            this.modelList.unshift(data);
                            this.current = data.Name;
                        }
                    }
                };
                this.issueItem = '';
                clear();
                o = load(data.name);
                vm.setTree(o);
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
