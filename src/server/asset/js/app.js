source = new EventSource('/events');

function AlertBox(value){
    var alertBox = '<div class="alert-box error"><span></span>'+value+'</div>';
    return alertBox;
}

source.addEventListener('update',function(e){
    var data = JSON.parse(e.data);
    //console.log(data.Name);
    clear();
    load(data.Name);
    //d1.insertAdjacentHTML('beforeend',AlertBox(data));
});
