<template>
    <sui-segment>
        <sui-input icon="search" placeholder="Search..."></sui-input>
        <sui-item-group divided >
            <sui-item :key="model.name" v-for="model in paginatedData">
                <sui-button compact basic v-on:click="load(model)">
                    <sui-item-image wrapped v-on:click="load(model)" size="tiny" v-show="model.img" :src="model.img">
                    </sui-item-image>
                </sui-button>
                <sui-item-content>
                    <sui-header :color="isActive(model)">
                        {{ model.name }}
                        <span><sui-icon size="small" name="delete" v-on:click="remove(model)"/></sui-icon></span>
                    </sui-header>
                    <sui-item-extra>
                        <sui-button-group>
                            <sui-button size="mini" icon="newspaper" v-on:click=""></sui-button>
                            <sui-button size="mini" icon="camera" v-on:click=""></sui-button>
                            <sui-button size="mini" icon="thumbtack" v-on:click=""></sui-button>
                        </sui-button-group>
                    </sui-item-extra>
                </sui-item-content>
            </sui-item>
        </sui-item-group>
        <sui-button-group attached="bottom">
            <sui-button @click="prevPage" :disabled="pageNumber === 0" icon="caret left" label-position="left">Prev</sui-button>
            <sui-button @click="nextPage" :disabled="pageNumber >= pageCount -1" icon="caret right" label-position="right" floated="right">Next</sui-button>
        </sui-button-group>
    </sui-segment>
</template>

<script>
export default {
    data(){
        return {
            pageNumber: 0
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
            console.log(obj);
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
            return Math.floor(l/s);
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
</style>
