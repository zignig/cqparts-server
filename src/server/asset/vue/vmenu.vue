<template>
    <sui-segment>
        <sui-menu fixed="left" compact vertical inverted v-show="visible">
            <sui-menu-item :key="model.name" v-for="model in modelList">
                <sui-image wrapped v-on:click="load(model)" size="tiny" v-show="model.img" :src="model.img"></sui-image>
                    {{ model.name }}
            </sui-menu-item>
        </sui-menu>
    </sui-segment>
</template>

<script>
export default {
    data(){
        return {
            pageNumber: 0,
            visible: true 
        }
    },
    props: {
        current: "",
        modelList: {
            type: Array,
        },
        size:{
            type: Number,
            required: false,
            default: 6,
        }
    },
    methods:{
        viz: function(){
            this.visible = !this.visible;
        },
        nextPage: function(){
            this.pageNumber++;
        },
        prevPage: function(){
            this.pageNumber--;
        },
        isActive: function(obj){
            if (this.current == obj.name){
                return "green"
            }
            else {
                return "black"
            }
        },
        load: function(obj){
            clear();
            load(obj.name);
            EventBus.$emit('select',obj);
        },
        remove : function(index){
            this.modelList.splice(this.modelList.indexOf(index),1);
            EventBus.$emit('delete item',index);
        }
    },
    computed:{
        pageCount(){
            let l = this.modelList.length,
                s = this.size;
            return Math.ceil(l/s);
        },
        paginatedData(){
            const start = this.pageNumber * this.size,
                end = start + this.size;
            return this.modelList
                .slice(start, end);
        }}
};

</script>

<style>
.nogap {
    padding-top: 0.0 !important;
    padding-right: 0.0 !important;
    padding-left: 0.0 !important;
    padding-bottom: 0.0 !important;
}
</style>
