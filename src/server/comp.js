

 var model = Vue.component('model',
{ template : '#model',
 name: 'model',

    props: {
        value: String, 
    },

    data: function(){
        return { models: ['one','two'] }
    },

    created: function(){
        this.models= [ 'one','two','three','case','train' ];
    },
    methods:{
        load: function(event){
            console.log(this.value);
        }
    
}}); 

