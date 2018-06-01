<template>
    <div class="pure-menu">
    <ul class="pure-menu-list">
        <li class="pure-menu-item" v-for="obj in models">
            <a class="pure-menu-link" v-on:click="load">{{ obj }}</a>
        </li>
    </ul>
    </div>
</template>

<script>
export default { 
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
    }
}

</script>

<style>
</style>
